package controllers

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
)

var NoMoreStdinInputError error = errors.New("Reached the end of the input or stdin was closed or an unexpected error happens")

func GetNumberFromStdin() (int, error) {
	return GetNumberFrom(os.Stdin)
}

func GetNumberFrom(reader io.Reader) (int, error) {
	var scanner = bufio.NewScanner(reader)
	for true {
		hasInput := scanner.Scan()
		if !hasInput {
			break
		}
		text := scanner.Text()
		number, err := strconv.Atoi(text)
		if err == nil {
			return number, nil
		} else {
			return 0, err
		}
	}
	return 0, NoMoreStdinInputError
}

func GetTextFromStdin() (string, error) {
	return GetTextFrom(os.Stdin)
}

func GetTextFrom(reader io.Reader) (string, error) {
	var scanner = bufio.NewScanner(reader)
	for true {
		hasInput := scanner.Scan()
		if !hasInput {
			break
		}
		text := scanner.Text()
		return text, nil
	}
	return "", NoMoreStdinInputError
}
