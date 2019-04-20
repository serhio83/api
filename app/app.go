package app

import (
	"log"

	"github.com/serhio83/api/pkg"
	"github.com/serhio83/api/pkg/config"
	"github.com/serhio83/api/pkg/mongo"
	"github.com/serhio83/api/pkg/server"
)

type App struct {
	server  *server.Server
	session *mongo.Session
	config  *root.Config
}

func(a *App) Initialize() {
	a.config = config.GetConfig()
	var err error
	a.session, err = mongo.NewSession(a.config.Mongo)
	if err != nil {
		log.Fatalln("unable to connect to mongodb")
	}

	u := mongo.NewUserService(a.session.Copy(), a.config.Mongo)
	a.server = server.NewServer(u, a.config)
}

func (a *App) Run() {
	defer a.session.Close()
	a.server.Start()
}
