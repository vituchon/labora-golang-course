package postgres

import (
	"github.com/vituchon/labora-golang-course/meeting-crud-api/model"
	"github.com/vituchon/labora-golang-course/meeting-crud-api/repositories"
)

type AnimalsRepository struct {
}

func NewAnimalsRepository() *AnimalsRepository {
	return &AnimalsRepository{}
}

// Gets all legal entities for a list of organizations
func (repo *AnimalsRepository) GetAll() ([]model.Animal, error) {
	rows, err := Conn.Query(`
		SELECT id, name, kind
		FROM animal`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	animals := []model.Animal{}
	for rows.Next() {
		animalPtr, err := scanAnimal(rows)
		if err != nil {
			return nil, err
		}
		animals = append(animals, *animalPtr)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return animals, nil
}

// Gets an specific legal animal by id
func (repo *AnimalsRepository) GetById(id int) (*model.Animal, error) {
	row := Conn.QueryRow(`
		SELECT id, name, kind
		FROM animal
		WHERE id = $1`, id)
	return scanAnimal(row)
}

// Creates a legal animal
func (repo *AnimalsRepository) Create(animal model.Animal) (*model.Animal, error) {
	createQuery := `INSERT INTO animal (name, kind) VALUES ($1, $2) returning id`
	err := Conn.QueryRow(createQuery, animal.Name, animal.Kind).Scan(&animal.Id)
	if err != nil {
		return nil, err
	}
	return &animal, nil
}

// Updates a legal animal (legal_id, session_id, name, is_default_biller)
func (repo *AnimalsRepository) Update(animal model.Animal) (*model.Animal, error) {
	if animal.Id == nil {
		return nil, repositories.EntityNotExistsErr
	}
	updateQuery := `UPDATE animal SET name = $1, kind = $2 WHERE id = $3`
	_, err := Conn.Exec(updateQuery, animal.Name, animal.Kind, animal.Id)
	if err != nil {
		return nil, err
	}
	return &animal, nil
}

func (repo *AnimalsRepository) Delete(id int) (err error) {
	deleteQuery := `DELETE FROM animal WHERE id = $1`
	_, err = Conn.Exec(deleteQuery, id)
	return
}

// Scans a row interpreting it as 'model.Animal' struct
func scanAnimal(rows RowScanner) (*model.Animal, error) {
	var animal model.Animal

	err := rows.Scan(&animal.Id, &animal.Name, &animal.Kind)
	if err != nil {
		return nil, err
	}
	return &animal, nil
}
