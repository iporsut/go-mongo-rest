package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HomeHandler ...
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	msg, _ := json.Marshal(map[string]string{"message": "Hello"})

	res.Write(msg)
}

// NotesHandler ...
func NotesHandler(db DbManager) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "application/json")

		j, err := json.Marshal(NotesResource{Notes: db.GetAll()})
		if err != nil {
			panic(err)
		}
		res.Write(j)
	})
}

func main() {
	db := NewMongoManager("localhost")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.Handle("/api/v1/notes", NotesHandler(db)).Methods("GET")
	http.Handle("/", r)

	log.Println("Listening on 8000")
	http.ListenAndServe(":8000", nil)
}
