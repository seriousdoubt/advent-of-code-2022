package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	const abc = "abcdefghijklmnopqrstuvwxyz"
	total := 0

	for scanner.Scan() {
		sack := scanner.Text()
		firstHalf := sack[0 : len(sack)/2]
		secondHalf := sack[len(sack)/2 : len(sack)]
		matchChar := ""
		matchFound := false
		points := 0

		fmt.Printf("First Half: %s\t Second Half: %s", firstHalf, secondHalf)

		for p1Char := 0; p1Char < len(sack)/2; p1Char++ {
			for p2Char := 0; p2Char < len(sack)/2; p2Char++ {
				//fmt.Printf("Evaluating %s to %s \n", string([]rune(firstHalf)[p1Char]), string([]rune(secondHalf)[p2Char]))
				if string([]rune(firstHalf)[p1Char]) == string([]rune(secondHalf)[p2Char]) {
					matchChar = string([]rune(firstHalf)[p1Char])
					fmt.Println("Match Found")
					matchFound = true
					if unicode.IsUpper([]rune(firstHalf)[p1Char]) {
						points = 26
					}
					points += strings.Index(abc, strings.ToLower(matchChar)) + 1
				}
				if matchFound == true {
					break
				}
			}
			if matchFound == true {
				break
			}
		}

		fmt.Printf("Line: %s \tPart1: %s \tPart2: %s \t Match: %s \tPoints: %d\n\n", sack, firstHalf, secondHalf, matchChar, points)
		total += points

		//find common

		//get val

	}

	fmt.Printf("Total Points: %d", total)
}
