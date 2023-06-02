package services

import (
	"errors"
	"fmt"

	"github.com/lib/pq"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/model"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/repositories/postgres"
)

type Bond struct {
	Person model.Person
	Animal model.Animal
}

func GetBondsByPersonIdsUsingRepositories(ids []int) ([]Bond, error) {
	bonds, err := bondsRepository.GetBondsOf(ids)
	if err != nil {
		errMsg := fmt.Sprintf("error while bonds animals : '%v'", err)
		return nil, errors.New(errMsg)
	}

	animalIds := make([]int, 0, len(bonds))
	personIds := make([]int, 0, len(bonds))
	for _, bond := range bonds {
		animalIds = append(animalIds, *&bond.AnimalId)
		personIds = append(personIds, *&bond.PersonId)
	}

	animals, err := animalsRepository.GetByIds(animalIds)
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

func GetBondsByPersonIdsDirectDBAccess(ids []int) ([]Bond, error) {
	rows, err := postgres.Conn.Query(`
		SELECT p.id,p.name,
					 a.id,a.name,a.kind
		FROM bond b
		INNER JOIN person p ON p.id = b.person_id
		INNER JOIN animal a ON a.id = b.animal_id
		WHERE b.person_id = ANY($1)`, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	bonds := []Bond{}
	for rows.Next() {
		bondPtr, err := scanBond(rows)
		if err != nil {
			return nil, err
		}
		bonds = append(bonds, *bondPtr)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return bonds, nil
}

// Scans a row interpreting it as 'services.Bond' struct
func scanBond(rows postgres.RowScanner) (*Bond, error) {
	var bond Bond

	err := rows.Scan(&bond.Person.Id, &bond.Person.Name, &bond.Animal.Id, &bond.Animal.Name, &bond.Animal.Kind)
	if err != nil {
		return nil, err
	}
	return &bond, nil
}
