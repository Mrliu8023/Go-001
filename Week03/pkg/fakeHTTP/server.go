package fakehttp

import (
	"context"
	"fmt"
	"net/http"
)

type fakeHandler struct {
}

func (fh *fakeHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Println("request from: ", req.RemoteAddr)
	return
}

func NewServer(addr string) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: &fakeHandler{},
	}
}

func Stop(s *http.Server, ctx context.Context) func() error {
	return func() error {
		return s.Shutdown(ctx)
	}
}
