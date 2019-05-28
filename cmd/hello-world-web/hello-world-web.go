package main

import (
	"fmt"
	"github.com/dherbrich/hello-world-pkg"
	"github.com/dherbrich/hello-world-web/internal/hellocontroller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("HELLO_PORT")

	router := mux.NewRouter()

	subrouterHello1 := router.PathPrefix("/hello1").Subrouter()
	subrouterHello1.Methods("GET").HandlerFunc(HandleGetHello1)

	subrouterHello2 := router.HandleFunc("/hello2", HandleGetHello2).Subrouter()
	subrouterHello2.Methods("GET")

	subrouterHello3 := router.HandleFunc("/hello3", hellocontroller.Greet).Subrouter()
	subrouterHello3.Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func HandleGetHello1(w http.ResponseWriter, _ *http.Request) {
	response := hello.SayHello("from 1")
	fmt.Fprintf(w, response)
}

func HandleGetHello2(w http.ResponseWriter, _ *http.Request) {
	response := hello.SayHello("from 2")
	fmt.Fprintf(w, response)
}
