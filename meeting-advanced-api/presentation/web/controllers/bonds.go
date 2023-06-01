package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/vituchon/labora-golang-course/meeting-advanced-api/model"
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

	bonds, err := DoGetBondsByPersonsIds(ids)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	WriteJsonResponse(response, http.StatusOK, bonds)
}

type BondDTO struct {
	Person model.Person
	Animal model.Animal
}

func DoGetBondsByPersonsIds(ids []int) ([]BondDTO, error) {

	bonds, err := bondsRepository.GetBondsOf(ids)
	if err != nil {
		errMsg := fmt.Sprintf("error while bonds animals : '%v'", err)
		return nil, errors.New(errMsg)
	}

	animalsIds := make([]int, 0, len(bonds))
	personIds := make([]int, 0, len(bonds))
	for _, bond := range bonds {
		animalsIds = append(animalsIds, *&bond.AnimalId)
		personIds = append(personIds, *&bond.PersonId)
	}

	animals, err := animalsRepository.GetByIds(animalsIds)
	if err != nil {
		errMsg := fmt.Sprintf("error while animals : '%v'", err)
		fmt.Println(errMsg)
		return nil, errors.New(errMsg)
	}

	persons, err := personRepository.GetByIds(personIds)
	if err != nil {
		errMsg := fmt.Sprintf("error while persons : '%v'", err)
		fmt.Println(errMsg)
		return nil, errors.New(errMsg)
	}

	var bondsDTO []BondDTO = make([]BondDTO, 0, len(bonds))
	for _, bond := range bonds {
		var personPtr *model.Person
		for _, person := range persons {
			if *person.Id == bond.PersonId {
				personPtr = &person
				break
			}
		}
		var animalPtr *model.Animal
		for _, animal := range animals {
			if *animal.Id == bond.AnimalId {
				animalPtr = &animal
				break
			}
		}

		if personPtr == nil || animalPtr == nil {
			errMsg := fmt.Sprintf("error while building bond(='%+v'), invalid reference to person or animal personPtr='%v' animalPtr='%v'", bond, personPtr, animalPtr)
			fmt.Println(errMsg)
		} else {
			var bondDTO BondDTO = BondDTO{
				Animal: *animalPtr,
				Person: *personPtr,
			}
			bondsDTO = append(bondsDTO, bondDTO)
		}
	}
	return bondsDTO, nil
}
