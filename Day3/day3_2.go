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
	group := make([]string, 0)
	points := 0
	totalPoints := 0

	for scanner.Scan() {

		if len(group) < 3 {
			group = append(group, scanner.Text())

			if len(group) == 3 {
				fmt.Printf("Group: %s\n", group)
				sack1 := group[0]
				sack2 := group[1]
				sack3 := group[2]
				match := ""

				for s1Char := 0; s1Char < len(sack1); s1Char++ {
					for s2Char := 0; s2Char < len(sack2); s2Char++ {
						for s3Char := 0; s3Char < len(sack3); s3Char++ {
							//fmt.Printf("s1Char: %c\t s2Char: %c\t s3Char: %c\n", sack1[s1Char], sack2[s2Char], sack3[s3Char])
							if sack1[s1Char] == sack2[s2Char] && sack1[s1Char] == sack3[s3Char] {
								points = 0
								match = string([]rune(sack1)[s1Char])
								if unicode.IsUpper([]rune(sack1)[s1Char]) {
									points = 26
								}
								points += strings.Index(abc, strings.ToLower(match)) + 1
								fmt.Printf("Found Match: %c\t for: %s\t points: %d\n", sack1[s1Char], group, points)
								totalPoints += points
								break
							}
						}
						if match != "" {
							break
						}
					}
					if match != "" {
						break
					}
				}
				group = nil
			}
		}
	}

	fmt.Printf("Total Points: %d", totalPoints)
}
