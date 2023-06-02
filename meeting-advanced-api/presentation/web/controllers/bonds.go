package controllers

import (
	"fmt"
	"net/http"

	"github.com/vituchon/labora-golang-course/meeting-advanced-api/presentation/services"
)

func GetBondsByPersonIds(response http.ResponseWriter, request *http.Request) {
	ids, err := ParseQueryParamAsInts(request, "ids")
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}

	bonds, err := services.GetBondsByPersonIds(ids)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	WriteJsonResponse(response, http.StatusOK, bonds)
}
