package main

import (
	"context"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestServerStarts(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("listen: %v", err)
	}

	mux := http.NewServeMux()
	srv := &http.Server{Handler: mux}

	done := make(chan error, 1)
	go func() {
		done <- srv.Serve(ln)
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		t.Fatalf("shutdown: %v", err)
	}

	if err := <-done; err != nil && err != http.ErrServerClosed {
		t.Fatalf("serve: %v", err)
	}
}
