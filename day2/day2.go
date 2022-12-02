package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"strings"
)

type Strategy struct {
	action, reaction string
}

func RewindFile(file *os.File) {
	_, err := file.Seek(0, io.SeekStart)
	if err != nil {
		fmt.Println(err)
	}
}

func CalculateScore(ActionPair Strategy) int {
	scoreMap := map[string]int {
		"rock": 1,
		"paper": 2,
		"scissors": 3,
	}
	score := 0
	// Initial score
	score += scoreMap[ActionPair.reaction]
	// Draw
	if ActionPair.reaction == ActionPair.action {
		score += 3
	// Win
	} else if ActionPair.reaction == "paper" && ActionPair.action == "rock" {
		score += 6
	} else if ActionPair.reaction == "scissors" && ActionPair.action == "paper" {
		score += 6
	} else if ActionPair.reaction == "rock" && ActionPair.action == "scissors" {
		score += 6
	}
	return score
}

func Task1(file *os.File) {
	RewindFile(file)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	totalScore := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fields := strings.Fields(line)
		var ActionPair Strategy
		switch fields[0] {
			case "A":
				ActionPair.action = "rock"
			case "B":
				ActionPair.action = "paper"
			case "C":
				ActionPair.action = "scissors"
		}
		switch fields[1] {
			case "X":
				ActionPair.reaction = "rock"
			case "Y":
				ActionPair.reaction = "paper"
			case "Z":
				ActionPair.reaction = "scissors"
		}
		totalScore += CalculateScore(ActionPair)
	}
	fmt.Println("Total score: ", totalScore)
}

func Task2(file *os.File) {
	RewindFile(file)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	totalScore := 0

	beatMap := map[string]string {
		"rock": "scissors",
		"paper": "rock",
		"scissors": "paper",
	}

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fields := strings.Fields(line)
		var ActionPair Strategy
		switch fields[0] {
			case "A":
				ActionPair.action = "rock"
			case "B":
				ActionPair.action = "paper"
			case "C":
				ActionPair.action = "scissors"
		}
		switch fields[1] {
			case "X":
				ActionPair.reaction = beatMap[ActionPair.action]
			case "Y":
				ActionPair.reaction = ActionPair.action
			case "Z":
				ActionPair.reaction = beatMap[beatMap[ActionPair.action]]
		}
		totalScore += CalculateScore(ActionPair)
	}
	fmt.Println("Total score: ", totalScore)
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