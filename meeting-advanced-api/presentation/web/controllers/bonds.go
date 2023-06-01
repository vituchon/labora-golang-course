package controllers

import (
	"fmt"
	"net/http"

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

func GetBondsByPersonsIds(response http.ResponseWriter, request *http.Request) {
	ids, err := ParseQueryParamAsInts(request, "ids")
	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
	}

	// DE ACA PARA ABAJO MOVER A SERVICES!
	bonds, err := bondsRepository.GetBondsOf(ids)
	if err != nil {
		errMsg := fmt.Sprintf("error while bonds animals : '%v'", err)
		fmt.Println(errMsg)
		http.Error(response, errMsg, http.StatusInternalServerError)
		return
	}

	// HACER TRABAJO DE CONVERTIR
	WriteJsonResponse(response, http.StatusOK, bonds)
}
