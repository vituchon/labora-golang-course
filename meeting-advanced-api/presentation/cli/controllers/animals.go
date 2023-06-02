package controllers

import (
	"fmt"
	"io"
	"strconv"

	"github.com/vituchon/labora-golang-course/meeting-advanced-api/model"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/repositories"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/repositories/postgres"
)

var animalsRepository repositories.Animals

func init() {
	// INYECCION DE DEPEDENCIA (Repositorio en memoria o en base de daatos)
	// memoria
	//animalsRepository = memory.NewAnimalsRepository()

	// base de datos (postgres)
	animalsRepository = postgres.NewAnimalsRepository()
}

func GetAnimals(input io.Reader, output io.Writer) {
	animals, err := animalsRepository.GetAll()
	if err != nil {
		errMsg := fmt.Sprintf("error while retrieving animals : '%v'", err)
		fmt.Println(errMsg)
		return
	}
	outputStr := fmt.Sprintf("%+v\n", animals)
	output.Write([]byte(outputStr))
}

func GetAnimalById(input io.Reader, output io.Writer) {
	output.Write([]byte("Ingrese ID: "))
	id, err := GetNumberFrom(input)
	if err != nil {
		errMsg := fmt.Sprintf("error getting id from input '%v'\n", err)
		fmt.Println(errMsg)
		return
	}

	animal, err := animalsRepository.GetById(id)
	if err != nil {
		errMsg := fmt.Sprintf("error while retrieving animal(id=%d): '%v'\n", id, err)
		fmt.Println(errMsg)
		return
	}
	outputStr := fmt.Sprintf("%+v\n", animal)
	output.Write([]byte(outputStr))
}

func CreateAnimal(input io.Reader, output io.Writer) {
	output.Write([]byte("Ingrese Nombre: "))
	enteredName, err := GetTextFrom(input)
	if err != nil {
		errMsg := fmt.Sprintf("error getting text from input '%v'\n", err)
		fmt.Println(errMsg)
		return
	}
	output.Write([]byte("Ingrese Kind [0-2]): "))
	enteredKind, err := GetNumberFrom(input)
	if err != nil {
		errMsg := fmt.Sprintf("error getting number from input '%v'\n", err)
		fmt.Println(errMsg)
		return
	}
	var enteredAnimal model.Animal = model.Animal{
		Name: enteredName,
		Kind: model.AnimalKind(enteredKind),
	} // TODO : VALIDATE Kind
	createdAnimal, err := animalsRepository.Create(enteredAnimal)
	if err != nil {
		errMsg := fmt.Sprintf("error while creating Animal: '%v'", err)
		fmt.Println(errMsg)
		return
	}
	outputStr := fmt.Sprintf("%+v\n", createdAnimal)
	output.Write([]byte(outputStr))
}

func UpdateAnimal(input io.Reader, output io.Writer) {
	// TODO : Deberia pedir primero el ID y validar que exista.... luego opcionalmente (decisión de diseño UI/UX, pero en una CLI en vez de una GUI	) si el ID es válido se puede recuperar el animal y mostrar los valores actuales (o proporcionar un mecanismo para no tener que volver a introducirlos...)
	output.Write([]byte("Ingrese Nombre: "))
	enteredName, err := GetTextFrom(input)
	if err != nil {
		errMsg := fmt.Sprintf("error getting text from input '%v'\n", err)
		fmt.Println(errMsg)
		return
	}

	output.Write([]byte("Ingrese Kind [0-2]): "))
	enteredKind, err := GetNumberFrom(input)
	if err != nil {
		errMsg := fmt.Sprintf("error getting number from input '%v'\n", err)
		fmt.Println(errMsg)
		return
	}

	output.Write([]byte("Ingrese ID: "))
	enteredId, err := GetNumberFrom(input)
	if err != nil {
		errMsg := fmt.Sprintf("error getting number from input '%v'\n", err)
		fmt.Println(errMsg)
		return
	}
	var enteredAnimal model.Animal = model.Animal{
		Id:   &enteredId,
		Name: enteredName,
		Kind: model.AnimalKind(enteredKind),
	} // TODO : VALIDATE Kind
	updatedAnimal, err := animalsRepository.Update(enteredAnimal)
	if err != nil {
		errMsg := fmt.Sprintf("error while updating Animal: '%v'", err)
		fmt.Println(errMsg)
		return
	}
	outputStr := fmt.Sprintf("%+v\n", updatedAnimal)
	output.Write([]byte(outputStr))
}

func DeleteAnimal(input io.Reader, output io.Writer) {
	output.Write([]byte("Ingrese ID: "))
	enteredId, err := GetNumberFrom(input)
	if err != nil {
		errMsg := fmt.Sprintf("error getting number from input '%v'\n", err)
		fmt.Println(errMsg)
		return
	}
	err = animalsRepository.Delete(enteredId)
	if err != nil {
		errMsg := fmt.Sprintf("error while deleting animal(id=%d): '%v'", enteredId, err)
		fmt.Println(errMsg)
		return
	}
	output.Write([]byte(strconv.Itoa(enteredId)))
}
