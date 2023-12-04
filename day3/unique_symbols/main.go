package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputFile = "./input.txt"

func getSymbols(f io.Reader) string {
	var output string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		output += filterSymbols(sc.Text())
	}

	return getUniqueSymbols(output)
}

func filterSymbols(s string) string {
	filterOut := "0123456789."
	for _, i := range filterOut {
		s = strings.ReplaceAll(s, string(i), "")
	}

	return s
}

func getUniqueSymbols(s string) string {
	var unique string

	for _, i := range s {
		if !strings.ContainsRune(unique, i) {
			unique += string(i)
		}
	}

	return unique
}

func main() {
	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	fmt.Println(getSymbols(f))
}
