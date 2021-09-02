package server

import (
	"fmt"
	"net/http"

	"github.com/mgdelacroix/lipsync/config"
)

func NewServer(cfg *config.Config, port int) (*http.Server, error) {
	mux := http.NewServeMux()

	mux.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		if cfg.Image != "" {
			http.ServeFile(w, r, cfg.Image)
			return
		}
		fmt.Fprintf(w, "no file configured")
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	return srv, nil
}
