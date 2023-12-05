package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const inputFile = "./input.txt"

func main() {
	var sum int

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		var numbers []int
		var winning []int
		var points int
		var match int

		card := strings.Split(sc.Text(), ":")
		values := strings.Split(card[1], "|")
		for _, i := range strings.Split(values[0], " ") {
			if i != "" {
				n, err := strconv.Atoi(i)
				if err != nil {
					log.Fatal(err)
				}
				winning = append(winning, n)
			}
		}
		for _, i := range strings.Split(values[1], " ") {
			if i != "" {
				n, err := strconv.Atoi(i)
				if err != nil {
					log.Fatal(err)
				}
				numbers = append(numbers, n)
			}
		}

		for _, i := range numbers {
			for _, j := range winning {
				if i == j {
					match++
				}
			}
		}

		if match == 0 {
			points = 0
		} else if match == 1 {
			points = 1
		} else if match == 2 {
			points = 2
		} else {
			points = int(math.Pow(2, float64(match-1)))
		}
		sum += points
	}

	fmt.Println(sum)
}
