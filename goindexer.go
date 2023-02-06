package main

import (
	"fileparser"
	"filewalker"
	"models"
	"os"
	"zincinject"
)

func main() {
	// Variable creation
	var filePaths []string
	var messages []models.EmailMessage
	var path string

	if len(os.Args) < 2 {
		panic("Please include the folder to index")
	} else {
		path = string(os.Args[1])
	}

	// Retrieve all files on the directory
	filePaths = filewalker.DirectoryWalk(path)

	// Parse the messages
	for _, path := range filePaths {
		var message = fileparser.FileParser(path)
		messages = append(messages, message)
	}

	file, error := os.Create("body.txt")
	if error != nil {
		panic(error)
	}
	defer file.Close()

	// Inject messages
	for _, message := range messages {
		zincinject.ZincInject(message, file)
	}
}
