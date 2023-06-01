package controllers

import (
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/repositories"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/repositories/postgres"
)

var personRepository repositories.Persons

func init() {
	personRepository = postgres.NewPersonsStorage()
}
