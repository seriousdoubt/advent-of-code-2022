package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func determineWinner(myHand string, theirHand string) string {

	result := ""

	if myHand == "X" {
		if theirHand == "A" {
			result = "tie"
		} else if theirHand == "B" {
			result = "them"
		} else if theirHand == "C" {
			result = "me"
		}
	} else if myHand == "Y" {
		if theirHand == "A" {
			result = "me"
		} else if theirHand == "B" {
			result = "tie"
		} else if theirHand == "C" {
			result = "them"
		}
	} else if myHand == "Z" {
		if theirHand == "A" {
			result = "them"
		} else if theirHand == "B" {
			result = "me"
		} else if theirHand == "C" {
			result = "tie"
		}
	}
	return result
}

func handScore(hand string) int {

	points := 0

	if hand == "A" || hand == "X" {
		points = 1
	} else if hand == "B" || hand == "Y" {
		points = 2
	} else if hand == "C" || hand == "Z" {
		points = 3
	}

	return points
}

func determineHand(theirHand string, outcome string) string {

	//X means you need to lose,
	//Y means you need to end the round in a draw,
	//and Z means you need to win.

	myHand := ""

	if theirHand == "A" {
		if outcome == "X" {
			myHand = "Z"
		} else if outcome == "Y" {
			myHand = "X"
		} else if outcome == "Z" {
			myHand = "Y"
		}
	} else if theirHand == "B" {
		if outcome == "X" {
			myHand = "X"
		} else if outcome == "Y" {
			myHand = "Y"
		} else if outcome == "Z" {
			myHand = "Z"
		}
	} else if theirHand == "C" {
		if outcome == "X" {
			myHand = "Y"
		} else if outcome == "Y" {
			myHand = "Z"
		} else if outcome == "Z" {
			myHand = "X"
		}
	}

	return myHand
}

func main() {
	file, err := os.Open("input")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	//First col is what opponent is going ot play
	//Second col is what you should play

	//First col: A = Rock, B = Paper, C = Scissors
	//Second col: X = Rock, Y = Paper, Z = Scissors

	//Scoring: The score for a single round is the score for the shape you selected
	//(1 for Rock, 2 for Paper, and 3 for Scissors) plus the score for the outcome of
	//the round (0 if you lost, 3 if the round was a draw, and 6 if you won).

	myTotalScore := 0
	theirTotalScore := 0

	for scanner.Scan() {
		hands := strings.Split(scanner.Text(), " ")
		outcome := hands[1]
		theirHand := hands[0]
		myHand := determineHand(theirHand, outcome)
		result := ""
		myHandPoints := 0
		theirHandPoints := 0

		result = determineWinner(myHand, theirHand)
		myHandPoints = handScore(myHand)
		theirHandPoints = handScore(theirHand)

		if result == "me" {
			myTotalScore += myHandPoints + 6
			theirTotalScore += theirHandPoints + 0
		} else if result == "them" {
			myTotalScore += myHandPoints + 0
			theirTotalScore += theirHandPoints + 6
		} else if result == "tie" {
			myTotalScore += myHandPoints + 3
			theirTotalScore += theirHandPoints + 3
		}

		//fmt.Printf("The Result: %s\n\n", result)
		//fmt.Printf("My Hand: %s\nTheir Hand: %s\n\n", myHand, theirHand)
	}

	fmt.Printf("!!TOTAL SCORE!!\n\nMy Points: %d\nTheir Points: %d\n\n", myTotalScore, theirTotalScore)
}
