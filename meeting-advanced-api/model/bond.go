package model

// A relation (link) between two animals (one of them beign a Person)
type Bond struct {
	Id       *int `json:"id"`
	PersonId int  `json:"personId"`
	AnimalId int  `json:"animalId"`
}
