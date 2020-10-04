package main;

import (
	"log"
	"net/http"
	"./routes"
	"./utils"
	"github.com/gorilla/mux"
)

func main() {

	utils.MigrateDB()

	route := mux.NewRouter()

	server := http.Server{

		Addr: ":3306",
		Handler: route,
	}

	log.Println("Running on port 3306")

	log.Println(server.ListenAndServe())
}