package postgres

import (
	"github.com/lib/pq"
	"github.com/vituchon/labora-golang-course/meeting-advanced-api/model"
)

type BondsRepository struct {
}

func NewBondsRepository() *BondsRepository {
	return &BondsRepository{}
}

func (repo *BondsRepository) GetBondsOf(personsId []int) ([]model.Bond, error) {
	rows, err := Conn.Query(`
		SELECT id, person_id, animal_id
		FROM bond
		WHERE person_id = ANY($1)`, pq.Array(personsId))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	bonds := []model.Bond{}
	for rows.Next() {
		bondPtr, err := scanBond(rows)
		if err != nil {
			return nil, err
		}
		bonds = append(bonds, *bondPtr)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return bonds, nil
}

// Scans a row interpreting it as 'model.Bond' struct
func scanBond(rows RowScanner) (*model.Bond, error) {
	var bond model.Bond

	err := rows.Scan(&bond.Id, &bond.PersonId, &bond.AnimalId)
	if err != nil {
		return nil, err
	}
	return &bond, nil
}
