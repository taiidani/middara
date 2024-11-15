package server

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/taiidani/middara/internal/cache"
)

type Server struct {
	cache     cache.Cache
	publicURL string
	port      string
	*http.Server
}

//go:embed templates
var templates embed.FS

// DevMode can be toggled to pull rendered files from the filesystem or the embedded FS.
var DevMode = os.Getenv("DEV") == "true"

func NewServer(cache cache.Cache) *Server {
	mux := http.NewServeMux()

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Required PORT environment variable not present")
	}

	publicURL := os.Getenv("PUBLIC_URL")
	if publicURL == "" {
		publicURL = "http://localhost:" + port
	}

	srv := &Server{
		Server: &http.Server{
			Addr:    fmt.Sprintf(":%s", port),
			Handler: mux,
		},
		publicURL: publicURL,
		port:      port,
		cache:     cache,
	}
	srv.addRoutes(mux)

	return srv
}

func (s *Server) addRoutes(mux *http.ServeMux) {
	mux.Handle("GET /", http.HandlerFunc(s.indexHandler))
	mux.Handle("GET /assets/", http.HandlerFunc(s.assetsHandler))
	mux.Handle("GET /game/{id}", http.HandlerFunc(s.gameHandler))
	mux.Handle("POST /game/{id}", http.HandlerFunc(s.saveGameHandler))
	mux.Handle("POST /game", http.HandlerFunc(s.newGameHandler))
}

func renderHtml(writer http.ResponseWriter, code int, file string, data any) {
	log := slog.With("name", file, "code", code)

	var t *template.Template
	var err error
	if DevMode {
		t, err = template.ParseGlob("internal/server/templates/**")
	} else {
		t, err = template.ParseFS(templates, "templates/**")
	}
	if err != nil {
		log.Error("Could not parse templates", "error", err)
		return
	}

	log.Debug("Rendering file", "dev", DevMode)
	writer.WriteHeader(code)
	err = t.ExecuteTemplate(writer, file, data)
	if err != nil {
		log.Error("Could not render template", "error", err)
	}
}
