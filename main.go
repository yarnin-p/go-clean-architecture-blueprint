package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"my-go-project-template/config"
	"my-go-project-template/server"
)

func main() {
	// Initial config
	if err := config.Init(); err != nil {
		log.Fatalf("%s", err.Error())
	}

	// Creat new app
	app := server.NewApp()

	// Start server
	if err := app.Run(viper.GetString("PORT")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
