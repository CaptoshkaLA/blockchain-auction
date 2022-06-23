package main

import (
	. "Auction/Controller/Router"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Println("=======START LISTENING=======")

	if len(os.Args) == 1 {
		log.Fatal("missing port number!")
	}

	port := os.Args[1]

	var allowedOrigins handlers.CORSOption = handlers.AllowedOrigins([]string{"*"})
	var allowedMethods handlers.CORSOption = handlers.AllowedMethods([]string{"GET", "POST"})

	var funcHandler func(http.Handler) http.Handler = handlers.CORS(allowedMethods, allowedOrigins)

	var router *mux.Router = NewRouter(port)
	var handler http.Handler = funcHandler(router)
	http.ListenAndServe(":"+port, handler)
}
