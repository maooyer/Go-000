package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func startHttpServer(ctx context.Context, addr string) error {
	fmt.Println("http server")
	s := http.Server{
		Addr: addr,
	}
	go func() {
		<-ctx.Done()
		fmt.Println("http shutdown")
		s.Shutdown(context.TODO())
	}()
	return s.ListenAndServe()
}

func registerSignal(ctx context.Context) error {
	quit := []os.Signal{os.Interrupt, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT}
	sig := make(chan os.Signal, len(quit))
	signal.Notify(sig, quit...)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("signal ctx done")
			return ctx.Err()
		case <-sig:
			fmt.Println("signal quit")
			return errors.New("signal quit")
		}
	}
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	// http server
	g.Go(func() error {
		return startHttpServer(ctx, ":9090")
	})

	// signal
	g.Go(func() error {
		return registerSignal(ctx)
	})

	err := g.Wait()
	fmt.Println("first error: ", err)
}
