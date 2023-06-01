package services

import (
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/repositories"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/repositories/postgres"
)

var animalsRepository repositories.Animals
var personRepository repositories.Persons
var bondsRepository repositories.Bonds

func init() {
	animalsRepository = postgres.NewAnimalsStorage()
	personRepository = postgres.NewPersonsStorage()
	bondsRepository = postgres.NewBondsStorage()
}
