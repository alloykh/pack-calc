package server

import (
	"context"
	"io"
	"log"
	"net/http"
	"time"

	stdLogger "github.com/alloykh/pack-calc/pkg/log"
)

type Server struct {
	name   string
	server *http.Server
}

func NewServer(name, addr string, handler http.Handler) (srv *Server, shutdown func()) {

	srv = &Server{
		name: name,
		server: &http.Server{
			Addr:         addr,
			Handler:      handler,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 30 * time.Second,
			IdleTimeout:  120 * time.Second,
		}}

	return srv, func() {
		if err := srv.Stop(); err != nil {
			log.Printf("error on server shutdown: %s\n", err.Error())
		}
	}
}

func (s *Server) Start() {

	go func() {
		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe err: %v\n", err)
		}
	}()

}

func (s *Server) Stop() error {
	if s.server == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		return err
	}

	log.Printf("%s server has stopped\n", s.name)

	return nil
}

func RunServer(log stdLogger.Logger, name, address string, handler http.Handler) func() {
	srv, shutdown := NewServer(name, address, handler)
	srv.Start()

	log.Infof("%s server started on %s\n", name, address)
	return shutdown
}

func CorsMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func HealthCheckHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if _, err := io.WriteString(w, "{ \"success\": true }"); err != nil {
			log.Printf("io.WriteString failed: %s", err)
		}
	})
}
