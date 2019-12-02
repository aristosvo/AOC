package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"strconv"
	"flag"
)

func main() {

	var filename = flag.String("file", "input.txt", "file to analyze")

	flag.Parse()
    file, err := os.Open(*filename)
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close()
	
	result:=0

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		weight, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		
		fuel := calculateFuel(weight)
		result += fuel
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
	}
	fmt.Println(result)
}

func calculateFuel(weight int) int {
	fuel :=  (weight / 3 ) - 2
    if fuel <= 0 {
        return 0
    }
    return fuel + calculateFuel(fuel)
}