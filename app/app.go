package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello!!!!!"))
}

func startApp() {
	fmt.Println("Postgre text search running on port 9911")
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":9911", nil))
}
