package repositories

import (
	"errors"

	"github.com/vituchon/labora-golang-course/meeting-crud-api/model"
)

var EntityNotExistsErr error = errors.New("Entity doesn't exists")
var DuplicatedEntityErr error = errors.New("Duplicated Entity")
var InvalidEntityStateErr error = errors.New("Entity state is invalid")

type Animals interface { // NOTAR queno llamo AnimalsRepository pues el nombre del paquete sirve de prefijo y al usar estar interfaz desde otro paquete queda `repositories.Games`, ver  @presentation\web\controllers\games.go#gamesRepository
	GetAll() ([]model.Animal, error)
	GetById(id int) (*model.Animal, error)
	Create(game model.Animal) (created *model.Animal, err error)
	Update(game model.Animal) (updated *model.Animal, err error)
	Delete(id int) error
}
