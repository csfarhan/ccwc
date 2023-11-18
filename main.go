package main

import (
	"bufio"
	"ccwc/helper"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Init flag variables
	var bytesFlag, linesFlag, wordsFlag, charsFlag bool
	flag.BoolVar(&bytesFlag, "c", false, "Count bytes")
	flag.BoolVar(&linesFlag, "l", false, "Count lines")
	flag.BoolVar(&wordsFlag, "w", false, "Count words")
	flag.BoolVar(&charsFlag, "m", false, "Count characters")
	flag.Parse()

	// Default to all flags if none are provided
	if !bytesFlag && !linesFlag && !wordsFlag && !charsFlag {
		bytesFlag, linesFlag, wordsFlag, charsFlag = true, true, true, true
	}

	// Get filename from command line
	fileName := ""
	args := flag.Args()
	if len(args) != 1 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			fileName = line
			break
		}
	} else {
		fileName = args[0]
	}

	// Obtain file info
	file, err := os.Stat(fileName)
	if err != nil {
		fmt.Println("File does not exist")
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
	// Print number of characters
	if charsFlag {
		fmt.Println("Characters in file:", helper.CharInFile(fileName))
	}
}
