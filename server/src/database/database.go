package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Object : The database structure for each object stored in the database
type Object struct {
	Side        string   `json:"side"`
	Position    int      `json:"position"`
	Item        string   `json:"item"`
	Container   string   `json:"container"`
	Description string   `json:"description"`
	Keywords    []string `json:"keywords"`
}

var objects []Object

// init : Initialise the database handler
func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln("Error getting current working directory: ", err)
	}

	dataFile, err := os.Open(fmt.Sprintf("%v/database/data.json", wd))
	if err != nil {
		log.Fatalln("Error opening the database JSON file: ", err)
	}

	byteData, err := ioutil.ReadAll(dataFile)
	if err != nil {
		log.Fatalln("Error reading from the database JSON file: ", err)
	}

	err = json.Unmarshal(byteData, &objects)
	if err != nil {
		log.Fatalln("Error unmarshalling database JSON file: ", err)
	}
}

// GetObject : Return information about an object based on a keyword search
func GetObject(keyword string) (*[]Object, error) {
	var foundItems []Object

	for _, item := range objects {
		if inArray(item.Keywords, keyword) {
			foundItems = append(foundItems, item)
		}
	}

	if len(foundItems) > 0 {
		return &foundItems, nil
	}

	return nil, errors.New("No objects found")
}

// inArray : Check if a user specified keyword matches any existing items in the database
func inArray(arrayToSearch []string, keyword string) bool {
	for _, word := range arrayToSearch {
		if word == keyword {
			return true
		}
	}

	return false
}
