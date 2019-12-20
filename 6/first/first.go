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

	result := 0
	for k := range m {
		result += calculateDistance(m, k)
	}

	fmt.Println(result)
}

func calculateDistance(m map[string]string, startingPoint string) int {
	if startingPoint == "COM" {
		return 0
	}
	return 1 + calculateDistance(m, m[startingPoint])
}
