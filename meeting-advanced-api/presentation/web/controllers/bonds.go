package controllers

import (
	"fmt"
	"net/http"

	"github.com/vituchon/labora-golang-course/meeting-advanced-api/presentation/services"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/repositories"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/repositories/postgres"
)

var bondsRepository repositories.Bonds

func init() {
	// INYECCION DE DEPEDENCIA (Repositorio en memoria o en base de daatos)
	// memoria
	//animalsRepository = memory.NewAnimalsStorage()

	// base de datos (postgres)
	bondsRepository = postgres.NewBondsStorage()
}

// TODO rename to GetBondsByPersonIds (without s)
func GetBondsByPersonsIds(response http.ResponseWriter, request *http.Request) {
	ids, err := ParseQueryParamAsInts(request, "ids")
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}

	bonds, err := services.GetBondsByPersonsIds(ids)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	WriteJsonResponse(response, http.StatusOK, bonds)
}
