package main

import (
	"log"
	"os"

	"github.com/mishramadhav/advent-of-code-2023/day1/internal/fileutils"
)

func main() {
	fileName := "input.txt"
	log.Printf("Reading file %s\n", fileName)
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("error while reading file - %v", err)
	}

	defer file.Close()

	lineReader := fileutils.NewBufferedFileReader(file)
	ans := processFile(lineReader)

	log.Printf("Answer is %d\n", ans)
}
