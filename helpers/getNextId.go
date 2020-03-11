package helpers

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"regexp"
	"strconv"
)

// GetNextID returns the next id to be stored
func GetNextID() (id int) {
	id = 0
	exPath, _ := filepath.Abs("./")
	path := exPath + "/data/"
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		name := f.Name()
		if name != "__init__" {
			match, _ := regexp.MatchString("[0-9]", name)
			if match {
				parsedValue, _ := strconv.Atoi(name)
				if parsedValue > id {
					id = parsedValue
				}
			}
		}
	}
	if id != 0 {
		id++
	}
	return
}
