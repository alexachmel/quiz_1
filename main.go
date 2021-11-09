package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var fileName = flag.String("file", "problems.json", "file with quiz questions and answers in JSON")

// Init the model for json data
type Quiz struct {
	Question string `json: "question"`
	Answer   string `json: "answer"`
}

func main() {
	// 1. Open JSON file
	file := openFile()
	defer file.Close()

	// 2. Read file's content
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed to read file body with error message: %s\n", err.Error())
		os.Exit(0)
	}

	// Create empty array
	var problems []Quiz

	// 3. Unmarshall the content to the struct
	err = json.Unmarshal(bytes, &problems)
	if err != nil {
		fmt.Printf("Failed to unmarshall the file content with error message: %s\n", err.Error())
		os.Exit(0)
	}

	// 4. Print questions
	fmt.Println("Please, answer the following questions:")
	for _, problem := range problems {
		fmt.Printf("%s\n", problem.Question)
		arr := strings.Split(problem.Question, "+")
		a, _ := strconv.Atoi(arr[0])
		b, _ := strconv.Atoi(arr[1])
		c := a + b
		fmt.Printf("%s vs. %s\n", strconv.Itoa(c), problem.Answer)
	}
}

func openFile() *os.File {
	// Show flag parsing and help command
	flag.Parse()
	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Printf("There is no such JSON file in the directory: %s\n", *fileName)
		os.Exit(0)
	}
	return file
}
