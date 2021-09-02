package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/mgdelacroix/lipsync/server"

	"github.com/urfave/cli/v2"
)

func serveAction(c *cli.Context) error {
	port := c.Int("port")
	srv, err := server.NewServer(port)
	if err != nil {
		return err
	}

	finished := make(chan bool, 1)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)

	go func() {
		fmt.Printf("starting server on port %d\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "error while running the server: %s\n", err)
		}
		finished <- true
	}()

	s := <-sigs
	fmt.Printf("Got signal %s, exiting...\n", s)
	srv.Close()

	<-finished
	return nil
}
