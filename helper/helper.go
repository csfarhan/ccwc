package helper

import (
	"bufio"
	"fmt"
	"os"
)

func LinesInFile(fileName string) []string {
	// Create io.Reader from file.
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// Create scanner from file.
	scanner := bufio.NewScanner(f)
	var result []string
	// Iterate over lines.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	// Length of result will be # of lines
	return result
}

func WordsInFile(filename string) int {
	// Create io.Reader from file.
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	// Create scanner from file
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	counter := 0
	for scanner.Scan() {
		counter++
	}

	return counter
}

func CharInFile(filename string) int {
	// Create io.Reader from file.
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	// Length of data will be # of characters
	return len(data)
}
