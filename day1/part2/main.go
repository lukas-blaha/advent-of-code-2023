package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFile = "./input.txt"

type Digits struct {
	newDigit bool
	first    struct {
		number int
		index  int
	}
	last struct {
		number int
		index  int
	}
}

func (d *Digits) findFirst(s, sub string, i int) {
	first := strings.Index(s, sub)
	if first != -1 {
		if d.newDigit {
			d.first.index = first
			d.first.number = i
		} else {
			if first < d.first.index {
				d.first.index = first
				d.first.number = i
			}
		}
	}
}

func (d *Digits) findLast(s, sub string, i int) {
	last := strings.LastIndex(s, sub)
	if last != -1 {
		if d.newDigit {
			d.last.index = last
			d.last.number = i
			d.newDigit = false
		} else {
			if last > d.last.index {
				d.last.index = last
				d.last.number = i
			}
		}
	}
}

func NewDigits(s string) *Digits {
	subs := []string{"1", "2", "3", "4",
		"5", "6", "7", "8", "9",
		"one", "two", "three",
		"four", "five", "six",
		"seven", "eight", "nine",
	}
	d := &Digits{
		newDigit: true,
	}

	for i, sub := range subs {
		var n int
		if i < 9 {
			n = i + 1
		} else {
			n = i - 8
		}
		d.findFirst(s, sub, n)
		d.findLast(s, sub, n)
	}

	return d
}

func (d *Digits) GetNumber() (int, error) {
	return strconv.Atoi(fmt.Sprintf("%d%d", d.first.number, d.last.number))
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
		d := NewDigits(sc.Text())
		n, err := d.GetNumber()
		if err != nil {
			log.Fatal(err)
		}
		sum += n
	}

	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
}
