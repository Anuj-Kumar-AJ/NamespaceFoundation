// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

package logging

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"namespacelabs.dev/go-ids"
)

const maxOutputToTerminal = 128

var Log = log.New(os.Stderr, "[grpclog] ", log.Ldate|log.Ltime|log.Lmicroseconds)

type interceptor struct{}

func (interceptor) unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	t := time.Now()
	ctx, reqid := logHeader(ctx, "request", info.FullMethod, req)
	resp, unaryErr := handler(ctx, req)
	err := AttachRequestIDToError(unaryErr, reqid)
	if err == nil {
		Log.Printf("%s: id=%s: took %v; response: %s", info.FullMethod, reqid, time.Since(t), serializeMessage(resp))
	} else {
		Log.Printf("%s: id=%s: took %v; error: %v", info.FullMethod, reqid, time.Since(t), err)
	}
	return resp, err
}

func (interceptor) streaming(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	t := time.Now()
	ctx, reqid := logHeader(stream.Context(), "stream", info.FullMethod, nil)
	streamErr := handler(serverStream{stream, ctx}, stream)
	err := AttachRequestIDToError(streamErr, reqid)
	if err == nil {
		Log.Printf("%s: id=%s: took %v, finished ok", info.FullMethod, reqid, time.Since(t))
	} else {
		Log.Printf("%s: id=%s: took %v; error: %v", info.FullMethod, reqid, time.Since(t), err)
	}
	return err
}

type serverStream struct {
	grpc.ServerStream
	ctx context.Context
}

func (s serverStream) Context() context.Context { return s.ctx }

func logHeader(ctx context.Context, what, fullMethod string, req interface{}) (context.Context, string) {
	// XXX establish request id propagation.
	reqid := ids.NewRandomBase32ID(16)
	peerAddr := "unknown"
	authType := "none"
	deadline := "none"
	if p, has := peer.FromContext(ctx); has {
		peerAddr = p.Addr.String()
		if p.AuthInfo != nil {
			authType = p.AuthInfo.AuthType()
		}
	}

	if t, ok := ctx.Deadline(); ok {
		left := time.Until(t)
		deadline = fmt.Sprintf("%v", left)
	}

	if req != nil {
		Log.Printf("%s: id=%s: request from %s (auth: %s, deadline: %s): %s", fullMethod, reqid, peerAddr, authType, deadline, serializeMessage(req))
	} else {
		Log.Printf("%s: id=%s: request from %s (auth: %s, deadline: %s)", fullMethod, reqid, peerAddr, authType, deadline)
	}

	return withRequestID(ctx, requestData{rid: reqid}), reqid
}

func Prepare(ctx context.Context, deps ExtensionDeps) error {
	var interceptor interceptor
	deps.Interceptors.ForServer(interceptor.unary, interceptor.streaming)
	return nil
}

func serializeMessage(msg interface{}) string {
	if msg == nil {
		return "<nil>"
	}

	reqBytes, _ := json.Marshal(msg)
	reqStr := string(reqBytes)
	if len(reqStr) > maxOutputToTerminal {
		return fmt.Sprintf("%s [...%d chars truncated]", reqStr[:maxOutputToTerminal], len(reqStr)-maxOutputToTerminal)
	}
	return reqStr
}
