package main

import (
	"flag"
	"fmt"
)

// Init the model for json data
/*
type Quiz struct {
    Question string 'json: "question"'
    Answer string 'json: "answer"'
}
*/

func main() {
	fileName := flag.String("file", "default.json", "file with quiz questions and answers in JSON")
	flag.Parse()
	_ = fileName


}

