package main

import (
	"log"
	"net/http"

	"github.com/erikschults/go-postgre-text-search/app/database"
	"github.com/erikschults/go-postgre-text-search/app/routes"
)

func startApp() {
	database.DBCon = database.New()
	database.Migrate()

	routes.RegisterRoutes()

	listen("9911")
}

func listen(port string) {
	done := make(chan bool)

	go http.ListenAndServe(":"+port, nil)
	log.Println("Postgre text search running on port " + port)

	<-done
}
