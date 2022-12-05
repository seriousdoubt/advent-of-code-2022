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

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func makeMove(quantity int, fromStack string, toStack string) (string, string) {
	crates := fromStack[0:quantity]
	crates = reverse(crates)
	fmt.Printf("quantity: %d\tfromstack: %s\ttostack: %s\n", quantity, fromStack, toStack)
	return fromStack[quantity:], crates + toStack
}

func strToInt(convString string) int {
	returnInt, err := strconv.Atoi(convString)
	check(err)

	return returnInt
}

func main() {

	stacks := make(map[int]string)
	// TOP -------------- BOTTOM
	stacks[1] = "ZVTBJGR"
	stacks[2] = "LVRJ"
	stacks[3] = "FQS"
	stacks[4] = "GQVFLNHZ"
	stacks[5] = "WMSCJTQR"
	stacks[6] = "FHCTWS"
	stacks[7] = "JNFVCZD"
	stacks[8] = "QFRWDZGL"
	stacks[9] = "PVWBJ"

	file, err := os.Open("moves.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//move 6 from 5 to 7
		//fmt.Println(scanner.Text())
		quan := strToInt(strings.Split(scanner.Text(), " ")[1])
		from := strToInt(strings.Split(scanner.Text(), " ")[3])
		to := strToInt(strings.Split(scanner.Text(), " ")[5])

		fmt.Print(scanner.Text())
		fmt.Printf("\nstack[%d]:  %s\tstack[%d]: %s\n", from, stacks[from], to, stacks[to])

		newFrom, newTo := makeMove(quan, stacks[from], stacks[to])

		fmt.Printf("Newtostack: %s\tNewfromstack: %s\n\n", newTo, newFrom)

		stacks[from] = newFrom
		stacks[to] = newTo
	}

	fmt.Print("Final Stacks\n\n")

	for i := 1; i <= 9; i++ {
		fmt.Printf("Stack %d\t %s\n", i, stacks[i])
	}

}
