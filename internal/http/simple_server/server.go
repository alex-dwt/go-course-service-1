package simple_server

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/pprof"
	"strconv"

	"alex/test/internal/http/handler"
	"alex/test/internal/service/user"
	"go.uber.org/zap"
)

type Server struct {
	server http.Server

	logger *zap.Logger
}

func New(
	port int,
	logger *zap.Logger,
	isDebug bool,
	userService *user.Service,
) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/get-example", handler.GetExampleRouteFunc(logger))
	mux.HandleFunc("/post-example", handler.PostExampleRoute)

	mux.HandleFunc("/user-create", func(w http.ResponseWriter, r *http.Request) {
		res, err := userService.CreateUserWithSomeLogic(context.Background(), "alex", rand.Intn(100))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, err.Error())
			return
		}

		io.WriteString(w, fmt.Sprintf("Result: %+v", res))
	})
	mux.HandleFunc("/user-update", func(w http.ResponseWriter, r *http.Request) {
		res, err := userService.GetUserWithSomeLogic(context.Background(), 4)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, err.Error())
			return
		}

		res.Name = "XXXXX"
		res.Age = uint8(99)

		if err := userService.SaveUser(context.Background(), res); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, err.Error())
			return
		}

		io.WriteString(w, fmt.Sprintf("SUCCESS"))
	})
	//mux.HandleFunc("/user-delete", func(w http.ResponseWriter, r *http.Request) {
	//
	//})
	//mux.HandleFunc("/user-get-with-age", func(w http.ResponseWriter, r *http.Request) {
	//
	//})

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
