package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/vituchon/labora-golang-course/meeting-advanced-api/model"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/repositories"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/repositories/postgres"
)

var animalsRepository repositories.Animals

func init() {
	// INYECCION DE DEPEDENCIA (Repositorio en memoria o en base de daatos)
	// memoria
	//animalsRepository = memory.NewAnimalsStorage()

	// base de datos (postgres)
	animalsRepository = postgres.NewAnimalsStorage()
}

func GetAnimals(response http.ResponseWriter, request *http.Request) {
	animals, err := animalsRepository.GetAll()
	if err != nil {
		errMsg := fmt.Sprintf("error while retrieving animals : '%v'", err)
		fmt.Println(errMsg)
		http.Error(response, errMsg, http.StatusInternalServerError)
		return
	}
	WriteJsonResponse(response, http.StatusOK, animals)
}

func GetAnimalById(response http.ResponseWriter, request *http.Request) {
	paramId := RouteParam(request, "id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		errMsg := fmt.Sprintf("Can not parse id from '%s'", paramId)
		fmt.Println(errMsg)
		http.Error(response, errMsg, http.StatusBadRequest)
		return
	}
	animal, err := animalsRepository.GetById(id)
	if err != nil {
		errMsg := fmt.Sprintf("error while retrieving animal(id=%d): '%v'\n", id, err)
		fmt.Println(errMsg)
		http.Error(response, errMsg, http.StatusInternalServerError)
		return
	}
	WriteJsonResponse(response, http.StatusOK, animal)
}

func CreateAnimal(response http.ResponseWriter, request *http.Request) {
	var animal model.Animal
	err := ParseJsonFromReader(request.Body, &animal)
	if err != nil {
		errMsg := fmt.Sprintf("error reading request body: '%v'", err)
		fmt.Println(errMsg)
		http.Error(response, errMsg, http.StatusBadRequest)
		return
	}

	created, err := animalsRepository.Create(animal)
	if err != nil {
		errMsg := fmt.Sprintf("error while creating Animal: '%v'", err)
		fmt.Println(errMsg)
		http.Error(response, errMsg, http.StatusInternalServerError)
		return
	}
	WriteJsonResponse(response, http.StatusOK, created)
}

func UpdateAnimal(response http.ResponseWriter, request *http.Request) {
	var animal model.Animal
	err := ParseJsonFromReader(request.Body, &animal)
	if err != nil {
		errMsg := fmt.Sprintf("error reading request body: '%v'", err)
		fmt.Println(errMsg)
		http.Error(response, errMsg, http.StatusBadRequest)
		return
	}
	updated, err := animalsRepository.Update(animal)
	if err != nil {
		errMsg := fmt.Sprintf("error while updating animal(id=%d): '%v'", animal.Id, err)
		fmt.Println(errMsg)
		http.Error(response, errMsg, http.StatusInternalServerError)
		return
	}
	WriteJsonResponse(response, http.StatusOK, updated)
}

func DeleteAnimal(response http.ResponseWriter, request *http.Request) {
	paramId := RouteParam(request, "id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		errMsg := fmt.Sprintf("Can not parse id from '%s'", paramId)
		fmt.Println(errMsg)
		http.Error(response, errMsg, http.StatusBadRequest)
		return
	}
	err = animalsRepository.Delete(id)
	if err != nil {
		errMsg := fmt.Sprintf("error while deleting animal(id=%d): '%v'", id, err)
		fmt.Println(errMsg)
		http.Error(response, errMsg, http.StatusInternalServerError)
		return
	}
	response.WriteHeader(http.StatusOK)
}
