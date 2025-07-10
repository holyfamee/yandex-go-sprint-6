package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	logger     *log.Logger
	httpServer *http.Server
}

func CreateServer(logger *log.Logger) *Server {
	r := chi.NewRouter()

	r.Get("/", handlers.ReturnData)
	r.Post("/upload", handlers.ReturnConverted)

	myServer := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger:     logger,
		httpServer: myServer,
	}
}
func (s *Server) Start() error {
	s.logger.Println("Сервер запускается на", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}
