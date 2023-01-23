package main

import (
	"advent/myutil"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("CalorieInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var elfCalories []int
	elfCalories = append(elfCalories, 0)
	for scanner.Scan() {
		if scanner.Text() != "" {
			currentCalories, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatal(err)
			}
			elfCalories[len(elfCalories)-1] = elfCalories[len(elfCalories)-1] + currentCalories
		} else {
			elfCalories = append(elfCalories, 0)
		}
	}

	_, max := myutil.MinMax(elfCalories)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(max)

	//Part 2
	first := 0
	second := 0
	third := 0
	for _, calorieTotal := range elfCalories {
		if calorieTotal > third {
			if calorieTotal > second {
				if calorieTotal > first {
					third = second
					second = first
					first = calorieTotal
				} else {
					third = second
					second = calorieTotal
				}
			} else {
				third = calorieTotal
			}
		}
	}
	part2Total := first + second + third
	fmt.Println(part2Total)

	//Or, with ReturnPositionMax and RemoveIndex, which I feel is easier to read but more resource intensive
	part2RunningTotal := 0
	for i := 0; i < 3; i++ {
		position, max := myutil.ReturnPositionMax(elfCalories)
		part2RunningTotal = part2RunningTotal + max
		elfCalories = myutil.RemoveIndex(elfCalories, position)
	}
	fmt.Println(part2RunningTotal)
}
