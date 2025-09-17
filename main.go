package main

import (
	"log"
	"net/http"
)



func main(){
	
	handler := http.NewServeMux()

	handler.HandleFunc("/hello", func(w http.ResponseWriter,r *http.Request){

		w.WriteHeader(200)

		w.Write([]byte("Hola World !"))
	})

	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}