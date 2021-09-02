package server

import (
	"fmt"
	"net/http"

	"github.com/mgdelacroix/lipsync/rss"

	"github.com/mgdelacroix/lipsync/config"
)

func NewServer(cfg *config.Config, port int) (*http.Server, error) {
	mux := http.NewServeMux()

	rssString, err := rss.GeneratePodcastRss(cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot generate rss feed: %w", err)
	}

	mux.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		if cfg.Image != "" {
			http.ServeFile(w, r, cfg.Image)
			return
		}
		fmt.Fprintf(w, "no file configured")
	})

	mux.HandleFunc("/feed.rss", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/rss+xml")
		fmt.Fprint(w, rssString)
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
