package internal

import (
	compute "combination-select/pkg"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Request struct {
	Weight int `json:"weight"`
	Items compute.Items `json:"items"`
}

func Handler(addr string){
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello World")
	})

	mux.HandleFunc("/combination-select", func(w http.ResponseWriter, r *http.Request) {
		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		var request Request
		json.Unmarshal(body, &request)

		result := compute.CombinationSelect(request.Weight, request.Items)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(result)

	})
	log.Fatal(http.ListenAndServe(addr, mux))
}