package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, directoryErr := filepath.Abs(filepath.Dir(os.Args[0]))
	files, filesError := ioutil.ReadDir(dir)
	if directoryErr != nil {
		log.Fatal(directoryErr)
	}

	if filesError != nil {
		log.Fatal(filesError)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}
