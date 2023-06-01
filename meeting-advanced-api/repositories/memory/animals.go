package memory

import (
	"sync"

	"github.com/vituchon/labora-golang-course/meeting-advanced-api/model"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/repositories"
)

type AnimalStorage struct {
	animalsById map[int]model.Animal
	idSequence  int
	mutex       sync.Mutex
}

func NewAnimalsStorage() *AnimalStorage {
	return &AnimalStorage{animalsById: make(map[int]model.Animal), idSequence: 0}
}

func (repo *AnimalStorage) GetAll() ([]model.Animal, error) {
	animals := make([]model.Animal, 0, len(repo.animalsById))
	for _, animal := range repo.animalsById {
		animals = append(animals, animal)
	}
	return animals, nil
}

func (repo *AnimalStorage) GetById(id int) (*model.Animal, error) {
	animal, exists := repo.animalsById[id]
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
	repo.animalsById[nextId] = animal
	repo.idSequence++ // can not reference idSequence as each update would increment all the animals Id by id (thus all will be the same)
	repo.mutex.Unlock()
	return &animal, nil
}

func (repo *AnimalStorage) Update(animal model.Animal) (updated *model.Animal, err error) {
	if animal.Id == nil {
		return nil, repositories.EntityNotExistsErr
	}
	repo.animalsById[*animal.Id] = animal
	return &animal, nil
}

func (repo *AnimalStorage) Delete(id int) error {
	delete(repo.animalsById, id)
	return nil
}
