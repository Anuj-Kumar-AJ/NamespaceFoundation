// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package deploy

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"github.com/kr/text"
	"namespacelabs.dev/foundation/engine/ops"
	"namespacelabs.dev/foundation/internal/console"
	"namespacelabs.dev/foundation/internal/console/renderwait"
	"namespacelabs.dev/foundation/runtime"
	"namespacelabs.dev/foundation/schema/orchestration"
	"namespacelabs.dev/foundation/std/planning"
	"namespacelabs.dev/foundation/workspace/tasks"
)

const (
	maxDeployWait      = 10 * time.Second
	tailLinesOnFailure = 10
)

func MaybeRenderBlock(env planning.Context, cluster runtime.ClusterNamespace, render bool) ops.WaitHandler {
	return func(ctx context.Context) (chan *orchestration.Event, func(context.Context, error) error) {
		if !render {
			return observeContainers(ctx, env, cluster, nil), func(ctx context.Context, err error) error { return err }
		}

		rwb := renderwait.NewBlock(ctx, "deploy")
		cleanup := func(ctx context.Context, waitErr error) error {
			// Make sure that rwb completes before further output below (for ordering purposes).
			if err := rwb.Wait(ctx); err != nil {
				if waitErr == nil {
					return err
				}
			}

			return waitErr
		}

		return observeContainers(ctx, env, cluster, rwb.Ch()), cleanup
	}
}

// observeContainers observes the deploy events (received from the returned channel) and updates the
// console through the `parent` channel.
func observeContainers(ctx context.Context, env planning.Context, cluster runtime.ClusterNamespace, parent chan *orchestration.Event) chan *orchestration.Event {
	ch := make(chan *orchestration.Event)
	t := time.NewTicker(maxDeployWait)
	startedDiagnosis := true // After the first tick, we tick twice as fast.

	go func() {
		if parent != nil {
			defer close(parent)
		}

		defer t.Stop()

		// Keep track of the pending ContainerWaitStatus per resource type.
		pending := map[string][]*runtime.ContainerWaitStatus{}
		helps := map[string]string{}

		runDiagnosis := func() {
			if startedDiagnosis {
				startedDiagnosis = false
				t.Reset(maxDeployWait / 2)
			}

			for resourceID, wslist := range pending {
				all := []*runtime.ContainerUnitWaitStatus{}
				for _, w := range wslist {
					all = append(all, w.Containers...)
					all = append(all, w.Initializers...)
				}
				if len(all) == 0 {
					// No containers are present yet, should still produce pod diagnostics.
					continue
				}

				buf := bytes.NewBuffer(nil)
				out := io.MultiWriter(buf, console.Debug(ctx))

				if help, ok := helps[resourceID]; ok && !env.Environment().GetEphemeral() {
					fmt.Fprintf(out, "For more information, run:\n  %s\n", help)
				}

				fmt.Fprintf(out, "Diagnostics retrieved at %s:\n", time.Now().Format("2006-01-02 15:04:05.000"))

				// XXX fetching diagnostics should not block forwarding events (above).
				for _, ws := range all {
					diagnostics, err := cluster.Cluster().FetchDiagnostics(ctx, ws.Reference)
					if err != nil {
						fmt.Fprintf(out, "Failed to retrieve diagnostics for %s: %v\n", ws.Reference.HumanReference, err)
						continue
					}

					fmt.Fprintf(out, "%s", ws.Reference.HumanReference)
					if diagnostics.RestartCount > 0 {
						fmt.Fprintf(out, " (restarted %d times)", diagnostics.RestartCount)
					}
					fmt.Fprintln(out, ":")

					switch {
					case diagnostics.Running:
						fmt.Fprintf(out, "  Running, logs (last %d lines):\n", tailLinesOnFailure)
						if err := cluster.Cluster().FetchLogsTo(ctx, text.NewIndentWriter(out, []byte("    ")), ws.Reference, runtime.FetchLogsOpts{TailLines: tailLinesOnFailure}); err != nil {
							fmt.Fprintf(out, "Failed to retrieve logs for %s: %v\n", ws.Reference.HumanReference, err)
						}

					case diagnostics.Waiting:
						fmt.Fprintf(out, "  Waiting: %s\n", diagnostics.WaitingReason)
						if diagnostics.Crashed {
							fmt.Fprintf(out, "  Crashed, logs (last %d lines):\n", tailLinesOnFailure)
							if err := cluster.Cluster().FetchLogsTo(ctx, text.NewIndentWriter(out, []byte("    ")), ws.Reference, runtime.FetchLogsOpts{TailLines: tailLinesOnFailure, FetchLastFailure: true}); err != nil {
								fmt.Fprintf(out, "Failed to retrieve logs for %s: %v\n", ws.Reference.HumanReference, err)
							}
						}

					case diagnostics.Terminated:
						if diagnostics.ExitCode > 0 {
							fmt.Fprintf(out, "  Failed: %s (exit code %d), logs (last %d lines):\n", diagnostics.TerminatedReason, diagnostics.ExitCode, tailLinesOnFailure)
							if err := cluster.Cluster().FetchLogsTo(ctx, text.NewIndentWriter(out, []byte("  ")), ws.Reference, runtime.FetchLogsOpts{TailLines: tailLinesOnFailure, FetchLastFailure: true}); err != nil {
								fmt.Fprintf(out, "Failed to retrieve logs for %s: %v\n", ws.Reference.HumanReference, err)
							}
						}
					}
				}

				if parent != nil {
					parent <- &orchestration.Event{
						ResourceId:  resourceID,
						WaitDetails: buf.String(),
					}
				} else {
					diagnostics, err := cluster.FetchEnvironmentDiagnostics(ctx)
					if err != nil {
						fmt.Fprintf(out, "Failed to retrieve environment diagnostics: %v\n", err)
					} else {
						_ = tasks.Attachments(ctx).AttachSerializable("diagnostics.json", "", diagnostics)
					}
				}
			}
		}

		for {
			select {
			case <-ctx.Done():
				return

			case ev, ok := <-ch:
				if !ok {
					return
				}

				if parent != nil {
					parent <- ev
				}

				if ev.Ready != orchestration.Event_READY {
					pending[ev.ResourceId] = nil
					helps[ev.ResourceId] = ev.RuntimeSpecificHelp

					failed := false
					for _, w := range ev.WaitStatus {
						cws := &runtime.ContainerWaitStatus{}
						if err := w.Opaque.UnmarshalTo(cws); err != nil {
							continue
						}
						pending[ev.ResourceId] = append(pending[ev.ResourceId], cws)

						var ctrs []*runtime.ContainerUnitWaitStatus
						ctrs = append(ctrs, cws.Containers...)
						ctrs = append(ctrs, cws.Initializers...)

						for _, ctr := range ctrs {
							if ctr.Status.Crashed || ctr.Status.Failed() {
								failed = true
							}
						}
					}

					if failed && !startedDiagnosis {
						runDiagnosis()
					}
				} else {
					delete(pending, ev.ResourceId)
				}

			case <-t.C:
				runDiagnosis()
			}
		}
	}()

	return ch
}
