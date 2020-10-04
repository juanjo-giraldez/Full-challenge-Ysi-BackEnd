package main;

import (
	"log"
	"net/http"
	"./routes"
	"./utils"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	utils.MigrateDB()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE"}, 
	})

	route := mux.NewRouter()

	routes.SetCompanyRoutes(route)

	server := http.Server{

		Addr: ":3306",
		Handler: c.Handler(route),
	}

	log.Println("Running on port 3306")

	log.Println(server.ListenAndServe())
}