package main

import (
	"crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/user", server.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/user", server.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", server.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", server.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/user/{id}", server.DeleteUser).Methods(http.MethodDelete)

	fmt.Println("Listen in the port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
