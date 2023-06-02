package memory

import (
	"sync"

	"github.com/vituchon/labora-golang-course/meeting-advanced-api/model"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/repositories"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/util"
)

type BondsRepository struct {
	bondsById  map[int]model.Bond
	idSequence int
	mutex      sync.Mutex
}

var bondsRepositoryInstance *BondsRepository = nil

func NewBondsRepository() *BondsRepository {
	if bondsRepositoryInstance == nil {
		bondsRepositoryInstance = &BondsRepository{bondsById: make(map[int]model.Bond), idSequence: 0}
	}
	return bondsRepositoryInstance
}

func (repo *BondsRepository) GetAll() ([]model.Bond, error) {
	bonds := make([]model.Bond, 0, len(repo.bondsById))
	for _, bond := range repo.bondsById {
		bonds = append(bonds, bond)
	}
	return bonds, nil
}

func (repo *BondsRepository) GetById(id int) (*model.Bond, error) {
	bond, exists := repo.bondsById[id]
	if !exists {
		return nil, repositories.EntityNotExistsErr
	}
	return &bond, nil
}

func (repo *BondsRepository) GetBondsOf(personsId []int) ([]model.Bond, error) {
	bonds := []model.Bond{}
	for _, bond := range repo.bondsById {
		if util.ContainsInt(personsId, *bond.Id) {
			bonds = append(bonds, bond)
		}
	}
	return bonds, nil
}

func (repo *BondsRepository) Create(bond model.Bond) (created *model.Bond, err error) {
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

func (repo *BondsRepository) Update(bond model.Bond) (updated *model.Bond, err error) {
	if bond.Id == nil {
		return nil, repositories.EntityNotExistsErr
	}
	repo.bondsById[*bond.Id] = bond
	return &bond, nil
}

func (repo *BondsRepository) Delete(id int) error {
	delete(repo.bondsById, id)
	return nil
}
