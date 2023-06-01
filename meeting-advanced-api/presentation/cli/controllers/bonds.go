package controllers

import (
	"fmt"
	"io"

	"github.com/vituchon/labora-golang-course/meeting-advanced-api/presentation/services"
)

func GetBondsByPersonsIds(input io.Reader, output io.Writer) {
	output.Write([]byte("Ingrese ID: "))
	id, err := GetNumberFrom(input)
	if err != nil {
		errMsg := fmt.Sprintf("error getting id from input '%v'\n", err)
		fmt.Println(errMsg)
		return
	}

	bonds, err := services.GetBondsByPersonsIds([]int{id})
	if err != nil {
		errMsg := fmt.Sprintf("error while retrieving bonds for person(id=%d): '%v'\n", id, err)
		fmt.Println(errMsg)
		return
	}
	outputStr := fmt.Sprintf("%+v\n", bonds)
	output.Write([]byte(outputStr))
}
