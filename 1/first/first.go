package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var filename = flag.String("file", "input.txt", "file to analyze")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		weight, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		fuel := calculate(weight)
		sum += fuel
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}

func calculate(input int) int {
	return (input / 3) - 2
}
