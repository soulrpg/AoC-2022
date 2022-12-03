package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"unicode"
)

func RewindFile(file *os.File) {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println(err)
	}
}

func GetItemPriority(item rune) int {
	if unicode.IsUpper(item) {
		return int(item) - 'A' + 27
	} else {
		return int(item) - 'a' + 1
	}
}

func Task1(file *os.File) {
	RewindFile(file)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	totalProrities := 0

	for fileScanner.Scan() {
		rucksack := fileScanner.Text()

		firstCompartment := rucksack[:len(rucksack) / 2]
		secondCompartment := rucksack[len(rucksack) / 2:]

		var itemFound rune

		// Should use a hashmap there
		for _, firstItem := range firstCompartment {
			for _, secondItem := range secondCompartment {
				if firstItem == secondItem {
					itemFound = firstItem
					break
				}
			}
			if itemFound != 0 {
				break
			}
		}

		totalProrities += GetItemPriority(itemFound)
	}
	fmt.Println("Total priority: ", totalProrities)
}

func Task2(file *os.File) {
	RewindFile(file)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	totalProrities := 0

	for {
		fileScanner.Scan()
		firstRucksack := fileScanner.Text()
		if len(firstRucksack) == 0 {
			break
		}
		fileScanner.Scan()
		secondRucksack := fileScanner.Text()
		fileScanner.Scan()
		thirdRucksack := fileScanner.Text()

		var itemFound rune

		// Should use a hashmap there
		for _, firstItem := range firstRucksack {
			for _, secondItem := range secondRucksack {
				for _, thirdItem := range thirdRucksack {
					if firstItem == secondItem && secondItem == thirdItem {
						itemFound = firstItem
						break
					}
				}
			}
			if itemFound != 0 {
				break
			}
		}

		totalProrities += GetItemPriority(itemFound)
	}
	fmt.Println("Total priority: ", totalProrities)
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