package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "./input.txt"
)

type Bag struct {
	Red   int
	Blue  int
	Green int
}

func main() {
	var sum int

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		bag := &Bag{}
		g := strings.Split(sc.Text(), ":")
		if err != nil {
			log.Fatal(err)
		}

		sets := strings.Split(g[1], ";")
		for _, set := range sets {
			balls := strings.Split(set, ",")

			for _, b := range balls {
				ball := strings.Split(b, " ")
				n, err := strconv.Atoi(ball[1])
				if err != nil {
					log.Fatal(err)
				}

				if ball[2] == "red" && n > bag.Red {
					bag.Red = n
				}
				if ball[2] == "blue" && n > bag.Blue {
					bag.Blue = n
				}
				if ball[2] == "green" && n > bag.Green {
					bag.Green = n
				}
			}
		}

		power := bag.Red * bag.Blue * bag.Green
		sum += power
	}

	fmt.Println(sum)
}
