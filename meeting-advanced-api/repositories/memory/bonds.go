package memory

import (
	"sync"

	"github.com/vituchon/labora-golang-course/meeting-advanced-api/model"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/repositories"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/util"
)

type BondsStorage struct {
	bondsById  map[int]model.Bond
	idSequence int
	mutex      sync.Mutex
}

func NewBondsStorage() *BondsStorage {
	return &BondsStorage{bondsById: make(map[int]model.Bond), idSequence: 0}
}

func (repo *BondsStorage) GetAll() ([]model.Bond, error) {
	bonds := make([]model.Bond, 0, len(repo.bondsById))
	for _, bond := range repo.bondsById {
		bonds = append(bonds, bond)
	}
	return bonds, nil
}

func (repo *BondsStorage) GetById(id int) (*model.Bond, error) {
	bond, exists := repo.bondsById[id]
	if !exists {
		return nil, repositories.EntityNotExistsErr
	}
	return &bond, nil
}

func (repo *BondsStorage) GetBondsOf(personsId []int) ([]model.Bond, error) {
	bonds := []model.Bond{}
	for _, bond := range repo.bondsById {
		if util.ContainsInt(personsId, *bond.Id) {
			bonds = append(bonds, bond)
		}
	}
	return bonds, nil
}

func (repo *BondsStorage) Create(bond model.Bond) (created *model.Bond, err error) {
	if bond.Id != nil {
		return nil, repositories.DuplicatedEntityErr
	}
	repo.mutex.Lock()
	nextId := repo.idSequence + 1
	bond.Id = &nextId
	repo.bondsById[nextId] = bond
	repo.idSequence++ // can not reference idSequence as each update would increment all the bonds Id by id (thus all will be the same)
	repo.mutex.Unlock()
	return &bond, nil
}

func (repo *BondsStorage) Update(bond model.Bond) (updated *model.Bond, err error) {
	if bond.Id == nil {
		return nil, repositories.EntityNotExistsErr
	}
	repo.bondsById[*bond.Id] = bond
	return &bond, nil
}

func (repo *BondsStorage) Delete(id int) error {
	delete(repo.bondsById, id)
	return nil
}
