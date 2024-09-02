package main

import (
	"context"
	"errors"
	"fmt"
	"go-http-server/server"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	log := log.New(w, "HTTP server: ", log.LstdFlags)

	server := server.GetServer(log)
	err := http.ListenAndServe("localhost:8080", server)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Fprintf(w, "Does this get called? %s\n", err)
	} else if err != nil {
		fmt.Fprintf(w, "Some other error. %s\n", err)
	}

	return err
}

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err)
		os.Exit(1)
	}
}
