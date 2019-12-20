package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var filename = flag.String("file", "input.txt", "file to analyze")
	flag.Parse()

	var width = 25
	var tall = 6
	file, err := os.Open(*filename)
	final := make([]string, 150)

	if err != nil {
		log.Fatalln("Couldn't open the file", err)
	}
	defer file.Close()
	b1 := make([]byte, width*tall)
	minimum := width * tall

	fmt.Println("Question 1:")
	fmt.Println(strings.Repeat("-", 30), "\nLast # is answer on the questioN:")

	for err == nil {
		n1, err := file.Read(b1)
		if err != nil {
			break
		}
		layer := string(b1[:n1])
		for k, t := range final {
			pos := string([]rune(layer)[k])
			if t == "" && (pos == "1" || pos == "0") {
				final[k] = pos
			}
		}
		zeroCount := strings.Count(layer, "0")
		if zeroCount < minimum {
			minimum = zeroCount
			twoos := strings.Count(layer, "2")
			ones := strings.Count(layer, "1")
			fmt.Println(ones * twoos)
		}
	}
	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("\nQuestion 2:", "\n"+strings.Repeat("-", 30), "\nDecypher the message below:")
	for i := 0; i < 6; i++ {
		input := strings.Join(final[i*25:i*25+25], "")
		result := strings.ReplaceAll(input, "0", " ")
		fmt.Println(result)
	}

}
