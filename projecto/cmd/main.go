package main

import (
	"fmt"
	"os"
)

func getFileNames(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var filenames []string = make([]string, 0, len(files))
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}
	return filenames, nil
}

func main() {
	var paths []string = []string{".", "..", "/home/vituchon"}

	var channel chan []string = make(chan []string)

	for i := 1; i <= len(paths); i++ {
		go func(i int) {
			path := paths[i-1]
			filenames, err := getFileNames(path)
			if err != nil {
				fmt.Println("Hubo un error:", err)
			}

			channel <- filenames
		}(i)
	}

	var allFilenames []string
	for i := 1; i <= len(paths); i++ {
		var filenames []string = <-channel
		allFilenames = append(allFilenames, filenames...)
	}
	fmt.Println("Me fui ", allFilenames)
}

/**

package main

import (
	"fmt"
	"os"
	"sync"
)

func getFileNames(path string) ([]string, error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var filenames []string = make([]string, 0, len(files))
	for _, file := range files {
		filenames = append(filenames, file.Name())
	}
	return filenames, nil
}

func main() {
	var wg sync.WaitGroup

	var paths []string = []string{".", ".."}

	wg.Add(len(paths))
	for i := 1; i <= len(paths); i++ {
		go func(i int) {
			path := paths[i-1]
			filenames, err := getFileNames(path)
			if err == nil {
				fmt.Printf("En path %s están %v\n", path, filenames)
			} else {
				fmt.Println("Hubo un error:", err)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Me fui")
}
*/
