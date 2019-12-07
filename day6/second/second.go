package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	var filename = flag.String("file", "../input.txt", "file to analyze")

	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	m := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		results := strings.Split(line, ")")
		m[results[1]] = results[0]
	}

	you := calculateDistance(m, "YOU")
	santa := calculateDistance(m, "SAN")
	result := 0
	for i, v := range you {
		if v != santa[i] {
			result = len(you[i:]) + len(santa[i:]) - 2
			break
		}
	}

	fmt.Println(result)
}

func calculateDistance(m map[string]string, startingPoint string) []string {
	if startingPoint == "COM" {
		var a []string
		return append(a, startingPoint)
	}
	return append(calculateDistance(m, m[startingPoint]), startingPoint)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
