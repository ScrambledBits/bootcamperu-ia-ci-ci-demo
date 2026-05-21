package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	addr = ":8080"
	root = "/srv"
)

func main() {
	check := flag.Bool("healthcheck", false, "probe http://127.0.0.1"+addr+"/ and exit 0 on 200")
	flag.Parse()

	if *check {
		runHealthcheck()
		return
	}

	runServer()
}

func runHealthcheck() {
	client := &http.Client{Timeout: 2 * time.Second}
	resp, err := client.Get("http://127.0.0.1" + addr + "/")
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		os.Exit(1)
	}
}

func runServer() {
	srv := &http.Server{
		Addr:              addr,
		Handler:           http.FileServer(http.Dir(root)),
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Printf("serving %s on %s", root, addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_ = srv.Shutdown(ctx)
}
