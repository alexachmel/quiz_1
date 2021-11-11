package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var fileName = flag.String("file", "problems.json", "file with quiz questions and answers in JSON")
var makeShuffle = flag.Bool("option", false, "whether you want to shuffle the quiz order each time it is run")

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
	r := rand.New(rand.NewSource(time.Now().Unix()))
	perm := r.Perm(len(problems))
	for curr_idx, perm_idx := range perm {
		var idx int
		if *makeShuffle {
			idx = perm_idx
		} else {
			idx = curr_idx
		}
		fmt.Printf("%s\n", problems[idx].Question)
		arr := strings.Split(problems[idx].Question, "+")
		a, _ := strconv.Atoi(arr[0])
		b, _ := strconv.Atoi(arr[1])
		fmt.Printf("%s vs. %s\n", strconv.Itoa(a+b), problems[idx].Answer)
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
