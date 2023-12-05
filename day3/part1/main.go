package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	inputFile     = "./input.txt"
	uniqueSymbols = "*$=#%/&+-@"
)

type Part struct {
	Number string
	Line   int
	First  int
	Last   int
	Start  bool
	Done   bool
}

type Symbol struct {
	Line  int
	Index int
}

func getSymbolIndex(s string, line int, symbols string) []Symbol {
	var indexes []Symbol

	for i, c := range s {
		for _, symbol := range symbols {
			if c == symbol {
				sym := Symbol{
					Line:  line,
					Index: i,
				}
				indexes = append(indexes, sym)
			}
		}
	}

	return indexes
}

func NewPart() *Part {
	return &Part{
		Start: false,
		Done:  false,
	}
}

func getPartIndex(s string, line int) []Part {
	part := NewPart()
	var parts []Part

	for i, c := range s {
		if n, err := strconv.Atoi(string(c)); err == nil {
			if !part.Start {
				part.Start = true
				part.Number += fmt.Sprint(n)
				part.First = i
				part.Line = line
			} else {
				part.Number += fmt.Sprint(n)
			}
			if i == len(s)-1 && part.Start {
				part.Last = i
				parts = append(parts, *part)
				part.Start = false
				part.Number = ""
			}
		} else {
			if part.Start {
				part.Last = i - 1
				part.Start = false
				part.Done = true
			}
			if part.Done {
				parts = append(parts, *part)
				part.Done = false
				part.Number = ""
			}
		}
	}

	return parts
}

func getCorrectParts(parts []Part, symbols []Symbol) []Part {
	var correctParts []Part

	for _, part := range parts {
		for _, symbol := range symbols {
			if symbol.Line == part.Line {
				if symbol.Index == part.First-1 || symbol.Index == part.Last+1 {
					correctParts = append(correctParts, part)
					break
				}
			} else if symbol.Line == part.Line-1 || symbol.Line == part.Line+1 {
				if symbol.Index+1 == part.First || symbol.Index-1 == part.Last {
					correctParts = append(correctParts, part)
					break
				} else if symbol.Index >= part.First && symbol.Index <= part.Last {
					correctParts = append(correctParts, part)
					break
				}
			}
		}
	}

	return correctParts
}

func main() {
	var line int
	var parts []Part
	var symbols []Symbol
	var sum int

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		for _, symbol := range getSymbolIndex(sc.Text(), line, uniqueSymbols) {
			symbols = append(symbols, symbol)
		}
		for _, part := range getPartIndex(sc.Text(), line) {
			parts = append(parts, part)
		}
		line++
	}

	correctParts := getCorrectParts(parts, symbols)
	for _, part := range correctParts {
		n, err := strconv.Atoi(part.Number)
		if err != nil {
			log.Fatal(err)
		}

		sum += n
	}

	fmt.Println(sum)
}
