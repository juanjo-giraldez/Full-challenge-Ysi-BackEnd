package routes

import (
	"../controllers"
	"github.com/gorilla/mux"
)

func SetCompanyRoutes( r *mux.Router) {

	subRouter := r.PathPrefix("/api").Subrouter()

	subRouter.HandleFunc("/companies", controllers.GetCompanies).Methods("GET")
	subRouter.HandleFunc("/companies", controllers.CreateCompany).Methods("POST")
	subRouter.HandleFunc("/companies/{id}", controllers.DeleteCompany).Methods("DELETE")

}