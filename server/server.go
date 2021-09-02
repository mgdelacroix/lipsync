package server

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/mgdelacroix/lipsync/rss"

	"github.com/mgdelacroix/lipsync/config"
)

var tmpl *template.Template

func init() {
	tmpl, _ = template.New("").Parse(`
<!doctype html>
<html>
  <head>
    <meta charset="utf-8" />
  </head>
  <body>
    <img width="400px" src="{{.PodcastURL}}/image" />
    <h1>{{.Title}}</h1>
    <p><a href="{{.PodcastURL}}/feed.rss">Podcast feed</a></p>
  </body>
</html>
`)
}

func NewServer(cfg *config.Config, port int) (*http.Server, error) {
	mux := http.NewServeMux()

	rssString, err := rss.GeneratePodcastRss(cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot generate rss feed: %w", err)
	}

	var templateBytes bytes.Buffer
	if err := tmpl.Execute(&templateBytes, cfg); err != nil {
		return nil, fmt.Errorf("cannot generate mainpage: %w", err)
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

	mux.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir(cfg.Directory))))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, templateBytes.String())
	})

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	return srv, nil
}
