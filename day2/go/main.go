package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getShapeScore(letter string) int {

	score := 0

	switch letter {
	case "X": // Rock
		score = 1
	case "Y": // Paper
		score = 2
	case "Z": // Scissors
		score = 3
	default:
		score = 0
	}

	return score
}

func getChallengeScore(letters string) int {

	score := 0

	// A - Rock
	// B - Paper
	// C - Scissors

	// X - My Rock
	// Y - My Paper
	// Z - My Scissors

	switch letters {
	case "A X", "B Y", "C Z":
		score = 3 // Draw
	case "A Y", "B Z", "C X":
		score = 6 // Win
	case "A Z", "B X", "C Y":
		score = 0 // Lose
	default:
		score = 0
	}

	return score

}

func calculateResultFromLine(fileLine string) int {

	totalScore := 0

	sls := strings.Split(fileLine, " ")

	shapeScore := getShapeScore(sls[1])
	challengeScore := getChallengeScore(fileLine)

	totalScore += shapeScore + challengeScore

	fmt.Println(fileLine+" =",
		strconv.Itoa(shapeScore)+" +",
		strconv.Itoa(challengeScore)+" =",
		strconv.Itoa(totalScore))

	return totalScore
}

func main() {

	result := 0

	file, err := os.Open("../data.txt")
	if err != nil {
		log.Fatal("Error while reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		text := scanner.Text()

		if text == "" {
			continue
		}

		result += calculateResultFromLine(text)
	}

	fmt.Println(result)
}
