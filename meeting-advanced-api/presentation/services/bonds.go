package services

import (
	"errors"
	"fmt"

	"github.com/vituchon/labora-golang-course/meeting-advanced-api/model"
)

type Bond struct {
	Person model.Person
	Animal model.Animal
}

func GetBondsByPersonsIds(ids []int) ([]Bond, error) {
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

	var bondsDTO []Bond = make([]Bond, 0, len(bonds))
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
			var bondDTO Bond = Bond{
				Animal: *animalPtr,
				Person: *personPtr,
			}
			bondsDTO = append(bondsDTO, bondDTO)
		}
	}
	return bondsDTO, nil
}
