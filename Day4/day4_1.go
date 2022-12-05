package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func nummyNums(pairOne string, pairTwo string) bool {
	p1n1, err := strconv.Atoi(strings.Split(pairOne, "-")[0])
	p1n2, err := strconv.Atoi(strings.Split(pairOne, "-")[1])
	p2n1, err := strconv.Atoi(strings.Split(pairTwo, "-")[0])
	p2n2, err := strconv.Atoi(strings.Split(pairTwo, "-")[1])

	check(err)

	if p1n1 >= p2n1 && p1n2 <= p2n2 {
		return true
	} else {
		return false
	}
}

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0

	for scanner.Scan() {
		pairOne := strings.Split(scanner.Text(), ",")[0]
		pairTwo := strings.Split(scanner.Text(), ",")[1]

		if nummyNums(pairOne, pairTwo) || nummyNums(pairTwo, pairOne) {
			fmt.Printf("This set contains an all emcompasing range: %s\n", scanner.Text())
			counter += 1
		}

	}

	fmt.Printf("Total encompasing ranges: %d", counter)

}
