package memory

import (
	"sync"

	"github.com/vituchon/labora-golang-course/meeting-crud-api/model"
	"github.com/vituchon/labora-golang-course/meeting-crud-api/repositories"
)

type AnimalStorage struct {
	gamesById  map[int]model.Animal
	idSequence int
	mutex      sync.Mutex
}

func NewAnimalsStorage() *AnimalStorage {
	return &AnimalStorage{gamesById: make(map[int]model.Animal), idSequence: 0}
}

func (repo *AnimalStorage) GetAll() ([]model.Animal, error) {
	games := make([]model.Animal, 0, len(repo.gamesById))
	for _, animal := range repo.gamesById {
		games = append(games, animal)
	}
	return games, nil
}

func (repo *AnimalStorage) GetById(id int) (*model.Animal, error) {
	animal, exists := repo.gamesById[id]
	if !exists {
		return nil, repositories.EntityNotExistsErr
	}
	return &animal, nil
}

func (repo *AnimalStorage) Create(animal model.Animal) (created *model.Animal, err error) {
	if animal.Id != nil {
		return nil, repositories.DuplicatedEntityErr
	}
	repo.mutex.Lock()
	nextId := repo.idSequence + 1
	animal.Id = &nextId
	repo.gamesById[nextId] = animal
	repo.idSequence++ // can not reference idSequence as each update would increment all the games Id by id (thus all will be the same)
	repo.mutex.Unlock()
	return &animal, nil
}

func (repo *AnimalStorage) Update(animal model.Animal) (updated *model.Animal, err error) {
	if animal.Id == nil {
		return nil, repositories.EntityNotExistsErr
	}
	repo.gamesById[*animal.Id] = animal
	return &animal, nil
}

func (repo *AnimalStorage) Delete(id int) error {
	delete(repo.gamesById, id)
	return nil
}
