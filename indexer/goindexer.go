package main

import (
	"fileparser"
	"filewalker"
	"fmt"
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
	fmt.Println("Retrieving files")
	filePaths = filewalker.DirectoryWalk(path)

	// Parse the messages
	fmt.Println("Parsing files")
	for _, path := range filePaths {
		var message = fileparser.FileParser(path)
		messages = append(messages, message)
	}

	// Creating error log file
	file, error := os.Create("errorlog.txt")
	if error != nil {
		panic(error)
	}
	defer file.Close()

	// Inject messages
	fmt.Println("Injecting files (errors will be saved in errorlog.txt)")
	for _, message := range messages {
		zincinject.ZincInject(message, file)
	}
}
