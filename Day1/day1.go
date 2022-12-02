package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func findHiVal(elfMap map[int]int) (int, int) {
	hiCals := 0
	hiNbr := 0
	for elfNbr, elfCals := range elfMap {
		if elfCals > hiCals {
			hiNbr = elfNbr
			hiCals = elfCals
		}
	}

	return hiNbr, hiCals
}

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	elfMap := map[int]int{}
	elfCounter := 0
	elfCals := 0

	for scanner.Scan() {
		//fmt.Println(scanner.Text())
		if scanner.Text() == "" {
			elfMap[elfCounter] = elfCals
			elfCounter += 1
			elfCals = 0
		} else {
			calsInt, err := strconv.Atoi(scanner.Text())
			check(err)
			elfCals += calsInt
		}
	}

	fmt.Print(elfMap)
	//find first
	hiNbr, hiVal := findHiVal(elfMap)
	delete(elfMap, hiNbr)
	fmt.Printf("\n\nElf with the most calories was %v with %v calories", hiNbr, hiVal)

	total := hiVal

	//find second
	hiNbr, hiVal = findHiVal(elfMap)
	delete(elfMap, hiNbr)
	fmt.Printf("\n\nElf with the most calories was %v with %v calories", hiNbr, hiVal)

	total += hiVal

	//find third
	hiNbr, hiVal = findHiVal(elfMap)
	delete(elfMap, hiNbr)
	fmt.Printf("\n\nElf with the most calories was %v with %v calories", hiNbr, hiVal)

	total += hiVal

	fmt.Printf("\n\n TOTAL top 3 cals: %v", total)
}
