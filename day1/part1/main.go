package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const inputFile = "./input.txt"

func getSum(n []string) int {
	var sum int

	for _, i := range n {
		s := string(i[0]) + string(i[len(i)-1])
		num, _ := strconv.Atoi(s)
		sum += num
	}

	return sum
}

func filterOutNumbers(fileName string) ([]string, error) {
	var n []string

	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		var line string

		for _, c := range sc.Text() {
			if _, err := strconv.Atoi(string(c)); err == nil {
				line += string(c)
			}
		}

		n = append(n, line)
	}

	if err = sc.Err(); err != nil {
		return nil, err
	}

	return n, nil
}

func main() {
	n, err := filterOutNumbers(inputFile)
	if err != nil {
		log.Panic(err)
	}

	sum := getSum(n)
	fmt.Println(sum)
}
