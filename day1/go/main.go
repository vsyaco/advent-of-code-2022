package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func minValueIndex(arr []int) (currentMin int, currentMinIndex int) {

	currentMin = math.MaxInt
	currentMinIndex = 0

	for i, v := range arr {
		if v <= currentMin {
			currentMin = v
			currentMinIndex = i
		}
	}

	if currentMin == math.MaxInt && currentMinIndex == 0 {
		currentMinIndex = 0
		currentMin = 0
	}

	return currentMin, currentMinIndex
}

func updateArray(arr []int, val int) []int {
	if len(arr) < 3 {
		arr = append(arr, val)
		return arr
	}

	minV, minI := minValueIndex(arr)

	if val > minV {
		arr[minI] = val
	}

	return arr
}

func sumArr(arr []int) int {
	summ := 0

	for _, v := range arr {
		summ += v
	}

	return summ
}

func main() {

	currentcal := 0
	maxcal := []int{0, 0, 0}

	file, err := os.Open("../data.txt")
	if err != nil {
		log.Fatal("Error while reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		text := scanner.Text()

		if text == "" {
			maxcal = updateArray(maxcal, currentcal)
			currentcal = 0
			continue
		}

		n, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println("Eror while converting", err)
			continue
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		currentcal += n
	}

	fmt.Println(sumArr(maxcal))
}
