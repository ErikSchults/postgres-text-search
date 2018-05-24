package routes

import (
	"encoding/json"
	"net/http"

	"github.com/erikschults/go-postgre-text-search/app/models"
	"github.com/gorilla/mux"
)

func RegisterRoutes() {
	r := mux.NewRouter()

	r.HandleFunc("/search", fullTextHandler).Methods("GET").Queries("txt", "{txt:[a-zA-Z0-9]+}")
	r.HandleFunc("/search", allBooksHandler).Methods("GET")

	http.Handle("/", r)
}

func fullTextHandler(w http.ResponseWriter, req *http.Request) {
	books, err := models.FullTextSearch(req.URL.Query().Get("txt"))
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	respondJSON(w, books)
}

func allBooksHandler(w http.ResponseWriter, req *http.Request) {
	books, err := models.AllBooks()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	respondJSON(w, books)
}

func respondJSON(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
