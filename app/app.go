package main

import (
	"fmt"
	"net/http"

	"github.com/erikschults/go-postgre-text-search/app/database"
	_ "github.com/lib/pq"
)

func index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello!!!!!"))
}

func startApp() {
	http.HandleFunc("/", index)

	database.New()
	listen("9911")
}

func listen(port string) {
	done := make(chan bool)

	go http.ListenAndServe(":"+port, nil)
	fmt.Println("Postgre text search running on port " + port)

	<-done
}
