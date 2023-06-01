package repositories

import (
	"errors"

	"github.com/vituchon/labora-golang-course/meeting-advanced-api/model"
)

var EntityNotExistsErr error = errors.New("Entity doesn't exists")
var DuplicatedEntityErr error = errors.New("Duplicated Entity")
var InvalidEntityStateErr error = errors.New("Entity state is invalid")

type Animals interface { // NOTAR queno llamo AnimalsRepository pues el nombre del paquete sirve de prefijo y al usar estar interfaz desde otro paquete queda `repositories.Animals`
	GetAll() ([]model.Animal, error)
	GetById(id int) (*model.Animal, error)
	GetByIds(ids []int) ([]model.Animal, error)
	Create(animal model.Animal) (created *model.Animal, err error)
	Update(animal model.Animal) (updated *model.Animal, err error)
	Delete(id int) error
}

type Bonds interface {
	GetBondsOf(personsId []int) ([]model.Bond, error)
}

type Persons interface {
	GetAll() ([]model.Person, error)
	GetById(id int) (*model.Person, error)
	GetByIds(ids []int) ([]model.Person, error)
	Create(person model.Person) (*model.Person, error)
	Update(person model.Person) (*model.Person, error)
	Delete(id int) (err error)
}
