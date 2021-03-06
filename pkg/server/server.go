package server

import (
	"log"
	"net/http"
	"os"

	"github.com/serhio83/api/pkg"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
  router *mux.Router
  config *root.ServerConfig
}

func NewServer(u root.UserService, config *root.Config) *Server {
  s := Server { 
    router: mux.NewRouter(),
    config: config.Server}
  
  a := authHelper{config.Auth.Secret}
  NewUserRouter(u, s.getSubrouter("/user"), &a)
  return &s
}

func(s *Server) Start() {
  log.Println("Listening on http://0.0.0.0:" + s.config.Port)
  if err := http.ListenAndServe(":" + s.config.Port, handlers.LoggingHandler(os.Stdout, s.router)); err != nil {
      log.Fatal("http.ListenAndServe: ", err)
  }
}

func(s *Server) getSubrouter(path string) *mux.Router {
  return s.router.PathPrefix(path).Subrouter()
}
