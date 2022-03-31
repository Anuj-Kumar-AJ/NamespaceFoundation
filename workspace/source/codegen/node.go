// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package codegen

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"cuelang.org/go/cue/format"
	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	"namespacelabs.dev/foundation/internal/engine/ops"
	"namespacelabs.dev/foundation/internal/engine/ops/defs"
	"namespacelabs.dev/foundation/internal/fnerrors"
	"namespacelabs.dev/foundation/internal/fnfs"
	"namespacelabs.dev/foundation/internal/uniquestrings"
	"namespacelabs.dev/foundation/languages"
	"namespacelabs.dev/foundation/schema"
	"namespacelabs.dev/foundation/workspace"
	"namespacelabs.dev/foundation/workspace/source/protos"
)

const wellKnownResource = ".foundation.std.types.Resource"

func Register() {
	ops.Register[*OpGenNode](generator{})
}

type generator struct{}

func (generator) Run(ctx context.Context, env ops.Environment, _ *schema.Definition, msg *OpGenNode) (*ops.DispatcherResult, error) {
	wenv, ok := env.(workspace.Packages)
	if !ok {
		return nil, errors.New("workspace.Packages required")
	}

	loc, err := wenv.Resolve(ctx, schema.PackageName(msg.Node.PackageName))
	if err != nil {
		return nil, err
	}

	return nil, generateNode(ctx, loc, msg.Node, msg.Protos, loc.Module.ReadWriteFS())
}

func ForNode(pkg *workspace.Package, available []*schema.Node) ([]*schema.Definition, error) {
	var allDefs []*schema.Definition

	if len(pkg.Provides) > 0 {
		var dl defs.DefList

		var list []*protos.FileDescriptorSetAndDeps
		for _, dl := range pkg.Provides {
			list = append(list, dl)
		}

		dl.Add("Generate Foundation exports", &OpGenNode{
			Node:   pkg.Node(),
			Protos: protos.Merge(list...),
		}, pkg.PackageName())

		if lst, err := dl.Serialize(); err != nil {
			return nil, err
		} else {
			allDefs = append(allDefs, lst...)
		}
	}

	for _, fmwk := range pkg.Node().GetCodegenFrameworks() {
		defs, err := languages.IntegrationFor(fmwk).GenerateNode(pkg, available)
		if err != nil {
			return nil, err
		}

		allDefs = append(allDefs, defs...)
	}

	return allDefs, nil
}

func generateNode(ctx context.Context, loc workspace.Location, n *schema.Node, parsed *protos.FileDescriptorSetAndDeps, fs fnfs.ReadWriteFS) error {
	var imports uniquestrings.List

	types := make([]*dpb.DescriptorProto, len(n.Provides))
	for k, p := range n.Provides {
		_, types[k] = protos.LookupDescriptorProto(parsed, p.GetType().GetTypename())
		if types[k] == nil {
			return fnerrors.InternalError("%s: type not found", p.GetType().GetTypename())
		}

		for _, f := range types[k].Field {
			if f.GetType() == dpb.FieldDescriptorProto_TYPE_MESSAGE {
				if f.GetTypeName() == wellKnownResource {
					imports.Add("namespacelabs.dev/foundation/std/fn:types")
				}
			}
		}
	}

	var out bytes.Buffer
	fmt.Fprintf(&out, "// This file is automatically generated.\npackage %s\n\n", packageName(loc))

	for _, imp := range imports.Strings() {
		fmt.Fprintf(&out, "import %q\n", imp)
	}

	fmt.Fprintf(&out, "#Exports: {\n")
	for k, p := range n.Provides {
		fmt.Fprintf(&out, "%s: {\n", p.Name)
		fmt.Fprintf(&out, "packageName: %q\n", n.PackageName)
		fmt.Fprintf(&out, "type: %q\n", p.Name)
		fmt.Fprintf(&out, "typeDefinition: ")
		enc := json.NewEncoder(&out)
		enc.SetIndent("", "  ")
		enc.Encode(p.Type)

		t := types[k]
		if len(t.GetField()) > 0 {
			fmt.Fprintf(&out, "with: {\n")
			var missingDef int
			for _, field := range t.GetField() {
				var t string
				if field.Type != nil {
					switch *field.Type {
					case dpb.FieldDescriptorProto_TYPE_DOUBLE,
						dpb.FieldDescriptorProto_TYPE_FLOAT:
						t = "float"
					case dpb.FieldDescriptorProto_TYPE_INT64,
						dpb.FieldDescriptorProto_TYPE_UINT64,
						dpb.FieldDescriptorProto_TYPE_INT32,
						dpb.FieldDescriptorProto_TYPE_FIXED64,
						dpb.FieldDescriptorProto_TYPE_FIXED32,
						dpb.FieldDescriptorProto_TYPE_UINT32,
						dpb.FieldDescriptorProto_TYPE_SFIXED32,
						dpb.FieldDescriptorProto_TYPE_SFIXED64,
						dpb.FieldDescriptorProto_TYPE_SINT32,
						dpb.FieldDescriptorProto_TYPE_SINT64:
						t = "int"
					case dpb.FieldDescriptorProto_TYPE_STRING:
						t = "string"
					case dpb.FieldDescriptorProto_TYPE_BYTES:
						t = "bytes"
					case dpb.FieldDescriptorProto_TYPE_ENUM:
						enum := protos.LookupEnumDescriptorProto(parsed, field.GetTypeName())
						var values []string
						for _, v := range enum.Value {
							values = append(values, fmt.Sprintf("%q", v.GetName()))
						}
						t = fmt.Sprintf("(%s)", strings.Join(values, "|"))
					case dpb.FieldDescriptorProto_TYPE_MESSAGE:
						if field.GetTypeName() == wellKnownResource {
							t = "types.#Resource"
						}
					}
				} else {
					t = "{...}"
				}

				if t == "" {
					missingDef++
				} else {
					if field.GetLabel() == dpb.FieldDescriptorProto_LABEL_REPEATED {
						t = "[..." + t + "]"
					}

					fmt.Fprintf(&out, "%s?: %s\n", field.GetJsonName(), t)
				}
			}
			if missingDef > 0 {
				fmt.Fprintf(&out, "...\n")
			}
			fmt.Fprintf(&out, "}\n")
		}

		fmt.Fprintf(&out, "}\n")
	}
	fmt.Fprintf(&out, "}\n")

	return fnfs.WriteWorkspaceFile(ctx, fs, loc.Rel("exports.fn.cue"), func(w io.Writer) error {
		formatted, err := format.Source(out.Bytes())
		if err != nil {
			return err
		}

		_, err = w.Write(formatted)
		return err
	})
}

func packageName(loc workspace.Location) string {
	if loc.Rel() == "." {
		return filepath.Base(loc.Module.ModuleName())
	}

	return filepath.Base(loc.Rel())
}
