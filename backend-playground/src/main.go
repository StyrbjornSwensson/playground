package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	answer := "Success"
	json.NewEncoder(w).Encode(answer)
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", test).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))

}
