package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello!!!"))
}

func startApp() {
	fmt.Println("Postgre text search running on port " + os.Getenv("PORT"))
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
