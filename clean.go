package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func Clean() {
	f, err := os.Open("words_alpha.txt")
	defer f.Close()
	if err != nil {
		fmt.Println("failed to open file")
		log.Fatal(err)
	}
	cleaned, err := os.Create("words.txt")
	defer cleaned.Close()
	if err != nil {
		fmt.Println("failed to create file")
		log.Fatal(err)
	}
	writer := bufio.NewWriter(cleaned)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := string(scanner.Bytes())
		word = strings.ToLower(word)
		if len(word) == 5 {
			writer.WriteString(word+"\n")
			cleaned.Sync()
		}
	}
	writer.Flush()
}
