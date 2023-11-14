package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var currentcal int = 0
var maxcal int = 0

func onElfChanged() {
	if currentcal >= maxcal {
		maxcal = currentcal
	}
	currentcal = 0
}

func main() {

	file, err := os.Open("../data.txt")
	if err != nil {
		log.Fatal("Error while reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		text := scanner.Text()

		if text == "" {
			onElfChanged()
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

	fmt.Println(maxcal)
}
