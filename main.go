package main

import (
	"ccwc/helper"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Init flag variables
	var bytesFlag, linesFlag, wordsFlag, charsFlag bool
	flag.BoolVar(&bytesFlag, "c", false, "")
	flag.BoolVar(&linesFlag, "l", false, "")
	flag.BoolVar(&wordsFlag, "w", false, "")
	flag.BoolVar(&charsFlag, "m", false, "")
	flag.Parse()

	// Get filename from command line
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Please provide a filename.")
		return
	}
	fileName := args[0]

	// Obtain file info
	file, err := os.Stat(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print file size
	if bytesFlag {
		fmt.Println("File size:", file.Size())
	}
	// Print number of lines
	if linesFlag {
		fmt.Println("Lines in file: ", len(helper.LinesInFile(fileName)))
	}
	// Print number of words
	if wordsFlag {
		fmt.Println("Words in file:", helper.WordsInFile(fileName))
	}

	if charsFlag {
		fmt.Println("Characters in file:", helper.CharInFile(fileName))
	}
}
