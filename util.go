package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

// Search - this is a function for searching a directory for a file/folder/both
func Search(pattern, pathforsearching, filetype string) []string {

	var files []string

	if filetype == "all" {
		filepath.Walk(pathforsearching, func(path string, f os.FileInfo, _ error) error {
			// if !f.IsDir() {
			r, err := regexp.MatchString(pattern, f.Name())
			if err == nil && r {
				absolutefilepath, err := filepath.Abs(path)
				if err != nil {
					log.Fatal(err)
				}
				files = append(files, absolutefilepath)
				// }
			}
			return nil
		})
	}

	if filetype == "file" {
		filepath.Walk(pathforsearching, func(path string, f os.FileInfo, _ error) error {
			if !f.IsDir() {
				r, err := regexp.MatchString(pattern, f.Name())
				if err == nil && r {
					absolutefilepath, err := filepath.Abs(path)
					if err != nil {
						log.Fatal(err)
					}
					files = append(files, absolutefilepath)
				}
			}
			return nil
		})
	}

	if filetype == "folder" {
		filepath.Walk(pathforsearching, func(path string, f os.FileInfo, _ error) error {
			if f.IsDir() {
				r, err := regexp.MatchString(pattern, f.Name())
				if err == nil && r {
					absolutefilepath, err := filepath.Abs(path)
					if err != nil {
						log.Fatal(err)
					}
					files = append(files, absolutefilepath)
				}
			}
			return nil
		})
	}

	return files
}

// CreateJSON - Takes local json data and merges into one var
func CreateJSON() {

	//List files
	files := Search(".json", "files/", "file")

	// Loop through files
	for _, element := range files {
		log.Println("Loading", element)

		// Read each file
		data, err := ioutil.ReadFile(element)
		if err != nil {
			log.Println("Error", err)
		}

		//Unmarshal to JSON
		singlegroup := new([]Person)
		err = json.Unmarshal(data, &singlegroup)
		if err != nil {
			log.Println("Error", err)
		}

		vara := *singlegroup
		// Merge with people struct (no overwrite)
		for _, v := range vara {
			people = append(people, v)
		}
	}
}
