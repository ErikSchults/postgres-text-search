package main

import (
	"log"
	"net/http"

	"github.com/erikschults/go-postgre-text-search/app/database"
)

func index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello!!!!!"))
}

func startApp() {
	database.DBCon = database.New()
	database.Migrate()

	http.HandleFunc("/", index)
	listen("9911")
}

func listen(port string) {
	done := make(chan bool)

	go http.ListenAndServe(":"+port, nil)
	log.Println("Postgre text search running on port " + port)

	<-done
}
