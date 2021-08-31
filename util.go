package main

import (
	"bufio"
	"os"
	"strings"
)

func readFile(filePath string) string {
	results := ""

	file, err := os.Open(filePath)
	if err != nil {
		println("Error reading the provided file path. Details: \n")
		println(err)

		return results
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		line := scanner.Text()

		if strings.TrimSpace((line)) != "" {
			results += line + "\n"
		}
	}

	if err := scanner.Err(); err != nil {
		println("Error whilst reading file from the provided file path. Details: \n")
		println(err)

		return results
	}

	// Remove last \n character
	results = strings.TrimRight(results, "\n")
	return results
}
