package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	//Size of input grid
	const xlen = 99
	const ylen = 99

	f, err := os.Open("ForestInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var input [][]byte
	for scanner.Scan() {
		input = append(input, []byte(scanner.Text()))
	}

	visibleTable := [ylen][xlen]bool{}

	for y := 0; y < len(input); y++ {
		var lineMax byte
		for x := 0; x < len(input[0]); x++ {
			if x == 0 {
				lineMax = input[y][x]
				visibleTable[y][x] = true
			} else {
				if input[y][x] > lineMax {
					lineMax = input[y][x]
					visibleTable[y][x] = true
				}
			}
		}
	}
	for x := 0; x < len(input[0]); x++ {
		var lineMax byte
		for y := 0; y < len(input); y++ {
			if x == 0 {
				lineMax = input[y][x]
				visibleTable[y][x] = true
			} else {
				if input[y][x] > lineMax {
					lineMax = input[y][x]
					visibleTable[y][x] = true
				}
			}
		}
	}

	for y := len(input) - 1; y >= 0; y-- {
		var lineMax byte
		for x := len(input[0]) - 1; x >= 0; x-- {
			if x == 0 {
				lineMax = input[y][x]
				visibleTable[y][x] = true
			} else {
				if input[y][x] > lineMax {
					lineMax = input[y][x]
					visibleTable[y][x] = true
				}
			}
		}
	}
	for x := len(input[0]) - 1; x >= 0; x-- {
		var lineMax byte
		for y := len(input) - 1; y >= 0; y-- {
			if x == 0 {
				lineMax = input[y][x]
				visibleTable[y][x] = true
			} else {
				if input[y][x] > lineMax {
					lineMax = input[y][x]
					visibleTable[y][x] = true
				}
			}
		}
	}

	visibleTotal := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			if visibleTable[y][x] {
				visibleTotal++
			}
		}
	}

	log.Println(visibleTotal)

	//Part 2
	//Assumptions: A tree of height 0 is still a tree, visibilty from the edges no longer matters
	maxScenicScore := 0
	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[0]); x++ {
			var iDist int
			var jDist int
			var kDist int
			var lDist int

			for i := 1; true; i++ {
				if x-i == -1 { //if we've hit gone past the edge of the forest, we need to stop
					break
				}
				//and we need to use the distance from the previous loop if we have gone past the edge
				iDist = i
				if input[y][x] <= input[y][x-i] {
					break
				}
			}

			for j := 1; true; j++ {
				if x+j == xlen {
					break
				}
				jDist = j
				if input[y][x] <= input[y][x+j] {
					break
				}
			}

			for k := 1; true; k++ {
				if y-k == -1 {
					break
				}
				kDist = k
				if input[y][x] <= input[y-k][x] {
					break
				}
			}

			for l := 1; true; l++ {
				if y+l == ylen {
					break
				}
				lDist = l
				if input[y][x] <= input[y+l][x] {
					break
				}
			}

			if iDist*jDist*kDist*lDist > maxScenicScore {
				maxScenicScore = iDist * jDist * kDist * lDist
			}

		}
	}
	log.Println(maxScenicScore)
}
