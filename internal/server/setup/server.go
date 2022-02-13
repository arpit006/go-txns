package setup

import (
	"github.com/arpit006/go_txns/internal/config"
	"github.com/arpit006/go_txns/internal/server"
	"log"
)

func StartServer() {
	// initialize
	if err := initialize(); err != nil {
		// TODO: fatal log
		log.Println("Error occurred while loading config. Error is ", err)
	}

	// dependencies

	// server
	router := server.NewRouter()
	apiServer := server.NewServer(router)
	apiServer.Serve(config.GetAddr())
}
