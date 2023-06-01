package postgres

import (
	"github.com/lib/pq"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/model"
)

type BondsStorage struct {
}

func NewBondsStorage() *BondsStorage {
	return &BondsStorage{}
}

func (repo *BondsStorage) GetBondsOf(personsId []int) ([]model.Bond, error) {
	rows, err := Conn.Query(`
		SELECT id, person_id, animal_id
		FROM bond
		WHERE person_id = ANY($1)`, pq.Array(personsId))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	relations := []model.Bond{}
	for rows.Next() {
		relationPtr, err := scanBond(rows)
		if err != nil {
			return nil, err
		}
		relations = append(relations, *relationPtr)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return relations, nil
}

// Scans a row interpreting it as 'model.Relation' struct
func scanBond(rows RowScanner) (*model.Bond, error) {
	var relation model.Bond

	err := rows.Scan(&relation.Id, &relation.PersonId, &relation.AnimalId)
	if err != nil {
		return nil, err
	}
	return &relation, nil
}
