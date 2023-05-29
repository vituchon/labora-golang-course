package model

import (
	"fmt"
	"strings"
)

type Animal struct {
	Id   *int       `json:"id"`
	Kind AnimalKind `json:"kind"`
	Name string     `json:"name"`
}

func (a Animal) String() string {
	var idField string = "\"Id\": nil "
	if a.Id != nil {
		idField = fmt.Sprintf("\"Id\": %d ", *a.Id)
	}

	var kindField = fmt.Sprintf("\"Kind\": %v ", a.Kind)
	var NameField = fmt.Sprintf("\"Name\": %v ", a.Name)

	str := strings.Join([]string{idField, kindField, NameField}, ",")
	return "{" + str + "}"
}

type AnimalKind int // Discuss: Podria ser Kind si el paquete fuera animal, pues quedaría animal.Kind al usars desde OTRO paquete... sin embargo como es model, creo que esta OK poner que el Kind se refiere a un animal. Pues el paquete model tiene modelizadas otra entidades que no están

// ¿acaso existe una visión unificada? o es mi apetencia de unidad.. exijencia de claridad y cohesión en las cosas que construimos... no lo sé.... así somos... abstracciones pensantes que crean otras abstracciones

const (
	CatKind    AnimalKind = 0
	DogKind               = 1
	TurtleKind            = 2
)

var animalKindNames = []string{
	CatKind:    "cat",
	DogKind:    "dog",
	TurtleKind: "turtle",
}

func (rk AnimalKind) String() string {
	if int(rk) < len(animalKindNames) {
		return animalKindNames[rk]
	}
	return fmt.Sprintf("String(): invalid value(=%d) for AnimalKind", rk)
}
