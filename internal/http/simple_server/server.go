package simple_server

import (
	"context"
	"errors"
	"net/http"
	"net/http/pprof"
	"strconv"

	"alex/test/internal/http/handler"
	"go.uber.org/zap"
)

type Server struct {
	server http.Server

	logger *zap.Logger
}

func New(port int, logger *zap.Logger, isDebug bool) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/get-example", handler.GetExampleRoute)
	mux.HandleFunc("/post-example", handler.PostExampleRoute)

	if isDebug {
		mux.HandleFunc("/debug/pprof/", pprof.Index)
	}

	r := Server{
		server: http.Server{
			Addr:    "localhost:" + strconv.Itoa(port),
			Handler: mux,
		},
		logger: logger,
	}
	return &r
}

func (s *Server) Start() error {
	if err := s.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	return s.server.Shutdown(context.Background())
}
