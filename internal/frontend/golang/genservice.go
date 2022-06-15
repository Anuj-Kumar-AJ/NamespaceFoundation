// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package golang

import (
	"context"
	"fmt"
	"os"
	"strings"
	"text/template"

	"namespacelabs.dev/foundation/internal/fnfs"
)

const (
	implFileName = "impl.go"
)

func CreateGolangScaffold(ctx context.Context, fsfs fnfs.ReadWriteFS, loc fnfs.Location) error {
	parts := strings.Split(loc.RelPath, string(os.PathSeparator))

	if len(parts) < 1 {
		return fmt.Errorf("Unable to determine package name.")
	}

	return generateGoSource(ctx, fsfs, loc.Rel(implFileName), serviceTmpl, serviceTmplOptions{
		Package: parts[len(parts)-1],
	})
}

type serviceTmplOptions struct {
	Package string
}

var serviceTmpl = template.Must(template.New(implFileName).Parse(`
package {{.Package}}

import (
	"context"

	"namespacelabs.dev/foundation/std/go/server"
)

type Service struct {
}

func (svc *Service) Echo(ctx context.Context, req *EchoRequest) (*EchoResponse, error) {
	// TODO add business logic here
	return EchoResponse{ Text: req.Text }
}

func WireService(ctx context.Context, srv server.Registrar, deps ServiceDeps) {
	svc := &Service{}
	proto.RegisterCountServiceServer(srv, svc)
}

`))
