package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Object : The database structure for each object stored in the database
type Object struct {
	Side        string   `json:"side"`
	Position    int      `json:"position"`
	Item        string   `json:"item"`
	Container   string   `json:"container"`
	Description string   `json:"description"`
	Frequent    bool     `json:"frequent"`
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

// compare : Helper function to perform all string checks in lowercase
func compare(str string, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}

// inArray : Check if a user specified keyword matches any existing items in the database
func inArray(arrayToSearch []string, keyword string) bool {
	for _, word := range arrayToSearch {
		if compare(word, keyword) {
			return true
		}
	}

	return false
}

// fuzzySearch : Search all JSON properties across all items for a keyword match
func fuzzySearch(item Object, keyword string) bool {
	found := false

	if compare(item.Side, keyword) {
		found = true
	}

	if compare(item.Item, keyword) {
		found = true
	}

	if compare(item.Container, keyword) {
		found = true
	}

	if compare(item.Description, keyword) {
		found = true
	}

	if inArray(item.Keywords, keyword) {
		found = true
	}

	return found
}

// GetObjects : Return a list containing information about all objects matching a keyword search
func GetObjects(keyword string) (*[]Object, error) {
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

// GetFuzzyObjects : Return information about all objects matching a fuzzy keyword search
func GetFuzzyObjects(keyword string) (*[]Object, error) {
	var foundItems []Object

	for _, item := range objects {
		if fuzzySearch(item, keyword) == true {
			foundItems = append(foundItems, item)
		}
	}

	if len(foundItems) > 0 {
		return &foundItems, nil
	}

	return nil, errors.New("No objects found")
}
