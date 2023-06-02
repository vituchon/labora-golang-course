package postgres

import (
	"github.com/lib/pq"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/model"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/repositories"
)

type PersonsRepository struct {
}

func NewPersonsRepository() *PersonsRepository {
	return &PersonsRepository{}
}

func (repo *PersonsRepository) GetAll() ([]model.Person, error) {
	rows, err := Conn.Query(`
		SELECT id, name
		FROM person`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	persons := []model.Person{}
	for rows.Next() {
		personPtr, err := scanPerson(rows)
		if err != nil {
			return nil, err
		}
		persons = append(persons, *personPtr)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return persons, nil
}

func (repo *PersonsRepository) GetById(id int) (*model.Person, error) {
	row := Conn.QueryRow(`
		SELECT id, name
		FROM person
		WHERE id = $1`, id)
	return scanPerson(row)
}

func (repo *PersonsRepository) GetByIds(ids []int) ([]model.Person, error) {
	rows, err := Conn.Query(`
		SELECT id, name
		FROM person
		WHERE id = ANY($1)`, pq.Array(ids))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	persons := []model.Person{}
	for rows.Next() {
		personPtr, err := scanPerson(rows)
		if err != nil {
			return nil, err
		}
		persons = append(persons, *personPtr)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return persons, nil
}

func (repo *PersonsRepository) Create(person model.Person) (*model.Person, error) {
	createQuery := `INSERT INTO person (name) VALUES ($1) returning id`
	err := Conn.QueryRow(createQuery, person.Name).Scan(&person.Id)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (repo *PersonsRepository) Update(person model.Person) (*model.Person, error) {
	if person.Id == nil {
		return nil, repositories.EntityNotExistsErr
	}
	updateQuery := `UPDATE person SET name = $1 WHERE id = $2`
	_, err := Conn.Exec(updateQuery, person.Name, person.Id)
	if err != nil {
		return nil, err
	}
	return &person, nil
}

func (repo *PersonsRepository) Delete(id int) (err error) {
	deleteQuery := `DELETE FROM person WHERE id = $1`
	_, err = Conn.Exec(deleteQuery, id)
	return
}

// Scans a row interpreting it as 'model.Person' struct
func scanPerson(rows RowScanner) (*model.Person, error) {
	var person model.Person

	err := rows.Scan(&person.Id, &person.Name)
	if err != nil {
		return nil, err
	}
	return &person, nil
}
