package main

import (
	"fmt"
	"os"
)

type ReadDirResult struct {
	err       error
	filenames []string
	path      string
}

func main() {
	getNotUsingChannels()
	fmt.Println("-------------")
	getUsingChannels()

}

func getFiles(path string) ([]os.DirEntry, error) {
	return os.ReadDir(path)
}

func getUsingChannels() {
	var paths []string = []string{".", "..", "/", "kas"}

	workersChannel := make(chan ReadDirResult)
	for workerNumber := 1; workerNumber <= len(paths); workerNumber++ {
		fmt.Printf("Starting go routine %d for path %s\n", workerNumber, paths[workerNumber-1])
		go func(workerNumber int, path string) {
			files, err := os.ReadDir(path)
			if err != nil {
				workersChannel <- ReadDirResult{err, nil, path}
				return
			}

			var filenames []string = make([]string, 0, len(files))
			for _, file := range files {
				filenames = append(filenames, file.Name())
			}
			fmt.Println(len(filenames))
			workersChannel <- ReadDirResult{nil, filenames, path}
		}(workerNumber, paths[workerNumber-1])
	}

	var readDirResults []ReadDirResult = make([]ReadDirResult, 0, len(paths))
	for i := 1; i <= len(paths); i++ {
		readDirResult := <-workersChannel
		readDirResults = append(readDirResults, readDirResult)
	}

	for index, readDirResult := range readDirResults {
		fmt.Printf("readDirResults[%d]=%+v\n", index, readDirResult)
	}
}

func getNotUsingChannels() {
	var paths []string = []string{".", "..", "/", "kas"}
	var readDirResults []ReadDirResult = make([]ReadDirResult, 0, len(paths))

	for i := 1; i <= len(paths); i++ {
		path := paths[i-1]

		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Printf("Error '%v'\n", err)
			continue
		}

		var filenames []string = make([]string, 0, len(files))
		for _, file := range files {
			filenames = append(filenames, file.Name())
		}
		readDirResult := ReadDirResult{nil, filenames, path}
		readDirResults = append(readDirResults, readDirResult)
	}

	for index, readDirResult := range readDirResults {
		fmt.Printf("readDirResults[%d]=%+v\n", index, readDirResult)
	}
}
