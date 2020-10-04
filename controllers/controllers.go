package controllers

import (

	"encoding/json"
	"fmt"
	"net/http"
	"../models"
	"../utils"
	"github.com/gorilla/mux"
	
)

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	
	companies := []models.Company{}

	db := utils.GetConnection()
	defer db.Close()
	db.Find(&companies)

	j, _ := json.Marshal(companies)
	utils.SendResponse(w, http.StatusOK, j)
}

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	
	company := models.Company{}

	db := utils.GetConnection()
	defer db.Close()
	
	err := json.NewDecoder(r.Body).Decode(&company)

	if err != nil {
		fmt.Println(err)
		utils.SendErr(w, http.StatusBadRequest)
		return
	}
	err = db.Create(&company).Error

	if err != nil {
		fmt.Println(err)
		utils.SendErr(w, http.StatusInternalServerError)
		return
	}


	j, _ := json.Marshal(company)
	utils.SendResponse(w, http.StatusCreated, j)
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	
	company := models.Company{}

	id := mux.Vars(r)["id"]

	db := utils.GetConnection()
	defer db.Close()

	db.Find(&company, id)

	if company.ID > 0 {
		db.Delete(company)
		utils.SendResponse(w, http.StatusOK, []byte(`{}`))

	} else {
		utils.SendErr(w, http.StatusNotFound)
	}
}