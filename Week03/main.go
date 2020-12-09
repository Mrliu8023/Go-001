package main

import (
	"context"
	"log"
	fakehttp "src/Go-001/Week03/pkg/fakeHTTP"
	"src/Go-001/Week03/pkg/signals"
	"sync"

	"golang.org/x/sync/errgroup"
)

var once *sync.Once

func main() {
	s1 := fakehttp.NewServer(":8001")
	s2 := fakehttp.NewServer(":8002")

	egs, _ := errgroup.WithContext(context.Background())
	egs.Go(s1.ListenAndServe)
	egs.Go(s2.ListenAndServe)

	stop := signals.SetupSignalHandler()

	go func() {
		<-stop
		egs.Go(fakehttp.Stop(s1, context.Background()))
		egs.Go(fakehttp.Stop(s2, context.Background()))
	}()

	if err := egs.Wait(); err != nil {
		// 不知道为什么退出会报  http: Server closed
		log.Fatal(err)
	}
}
