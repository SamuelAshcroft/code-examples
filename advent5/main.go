package main

import (
	"advent/myutil"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var PART1 bool = true

func main() {
	stackInput, err := os.Open("StackInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer stackInput.Close()

	//Place input into a variable
	var input [][]byte
	scanner := bufio.NewScanner(stackInput)
	for scanner.Scan() {
		var currentLine []byte
		lineLength := len(scanner.Text())
		for i := 1; i < lineLength; i = i + 4 {
			currentLine = append(currentLine, scanner.Text()[i])
		}
		input = append(input, currentLine)
	}

	//Rearrange input to represent the stacks of crates, transposed from input
	stacks := make([][]byte, len(input[0]))
	for i := 1; i <= len(input); i++ {
		line := input[len(input)-i]
		for j, v := range line {
			if string(v) != " " {
				stacks[j] = append(stacks[j], v)
			}
		}
	}

	moveInput, err := os.Open("MoveInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer moveInput.Close()

	scanner = bufio.NewScanner(moveInput)
	for scanner.Scan() {
		moveCrate(scanner.Text(), stacks)
	}

	var topcrates []byte
	for _, stack := range stacks {
		topcrate, _ := myutil.PopSlice(stack)
		topcrates = append(topcrates, topcrate)
	}
	fmt.Println(string(topcrates))
}

func moveCrate(moveCommandInput string, stacks [][]byte) {
	moveCommand := strings.Split(moveCommandInput, " ")
	moveAmount, err := strconv.Atoi(moveCommand[1])
	if err != nil {
		log.Println(err)
	}

	stack1Position, err := strconv.Atoi(moveCommand[3])
	if err != nil {
		log.Println(err)
	}
	stack1Position = stack1Position - 1
	stack1 := stacks[stack1Position]

	stack2Position, err := strconv.Atoi(moveCommand[5])
	if err != nil {
		log.Println(err)
	}
	stack2Position = stack2Position - 1
	stack2 := stacks[stack2Position]

	if PART1 {
		for i := 0; i < moveAmount; i++ {
			var crate byte
			crate, stack1 = myutil.PopSlice(stack1)
			stack2 = append(stack2, crate)
		}
	} else {
		tempSlice := stack1[len(stack1)-moveAmount:]
		stack1 = stack1[:len(stack1)-moveAmount]
		stack2 = append(stack2, tempSlice...)
	}

	stacks[stack1Position] = stack1
	stacks[stack2Position] = stack2
}
