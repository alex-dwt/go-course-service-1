package simple_server

import (
	"net/http"
	"strconv"

	"alex/test/internal/http/handler"
	"go.uber.org/zap"
)

type Server struct {
	mux  *http.ServeMux
	port int

	logger *zap.Logger
}

func New(port int, logger *zap.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/get-example", handler.GetExampleRoute)
	mux.HandleFunc("/post-example", handler.PostExampleRoute)

	r := Server{
		mux:    mux,
		port:   port,
		logger: logger,
	}

	return &r
}

func (s *Server) Start() error {
	return http.ListenAndServe(":"+strconv.Itoa(s.port), s.mux)
}
