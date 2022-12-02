package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"io"
)

func RewindFile(file *os.File) {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println(err)
	}
}

func TestMax(value int, currentMax *int) {
	if value > *currentMax {
		*currentMax = value
	}
}

// currentHighest values are sorted
func TestMaxMulti(value int, currentHighest *[]int) {
	for i, current := range *currentHighest {
		if value > current {
			// Add new element at the end of the array - value is not relevant
			*currentHighest = append(*currentHighest, 0)
			copy((*currentHighest)[i + 1:], (*currentHighest)[i:])
			(*currentHighest)[i] = value
			*currentHighest = (*currentHighest)[:len(*currentHighest) - 1]
			break
		}
	}
}

func SumSlice(inputSlice []int) int {
	sum := 0
	for _, value := range inputSlice {
		sum += value
	}
	return sum
}

func Task1(file *os.File) {
	RewindFile(file)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	caloriesSummed := make([]int, 1)
	caloriesSummed[0] = 0

	maxCalories := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			calorie, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
			}
			caloriesSummed[len(caloriesSummed) - 1] += calorie
		} else {
			TestMax(caloriesSummed[len(caloriesSummed) - 1], &maxCalories)
			caloriesSummed = append(caloriesSummed, 0)
		}
	}
	TestMax(caloriesSummed[len(caloriesSummed) - 1], &maxCalories)
	fmt.Println("Max calories carried: ", maxCalories)
}

func Task2(file *os.File) {
	RewindFile(file)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	caloriesSummed := make([]int, 1)
	caloriesSummed[0] = 0

	maxCalories := []int{0, 0, 0}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line != "" {
			calorie, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
			}
			caloriesSummed[len(caloriesSummed) - 1] += calorie
		} else {
			TestMaxMulti(caloriesSummed[len(caloriesSummed) - 1], &maxCalories)
			caloriesSummed = append(caloriesSummed, 0)
		}
	}
	TestMaxMulti(caloriesSummed[len(caloriesSummed) - 1], &maxCalories)
	fmt.Println("Max calories carried: ", SumSlice(maxCalories))
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	Task1(file)
	Task2(file)

	file.Close()
}