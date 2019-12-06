package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var filename = flag.String("file", "../input.txt", "file to analyze")
	var inputValue = flag.Int("value", 1, "input value")
	flag.Parse()

	csvfile, err := os.Open(*filename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()
	inputRecords, err := csv.NewReader(csvfile).Read()
	if err != nil {
		log.Fatalln("You did something stupid", err)
	}

	records, err := stringValuesToInts(inputRecords)
	if err != nil {
		log.Fatalln("You did something stupid", err)
	}

	IntCodeComputer(records, *inputValue)
}

// IntCodeComputer yeah
func IntCodeComputer(records []int, input int) {

	noRecords := len(records)

	n := 0
	for n < noRecords {
		opcode := records[n] % 100
		//probably not used for now
		mode3 := records[n] / 10000
		mode2 := (records[n] - mode3*10000) / 1000
		mode1 := (records[n] - mode2*1000) / 100
		if opcode == 99 {
			return
		}
		// fmt.Println("Opcode:", opcode)
		switch opcode {
		case 1:
			targetPosition := records[n+3]
			inputValue1 := calcInputValue(records, mode1, n+1)
			inputValue2 := calcInputValue(records, mode2, n+2)
			// fmt.Println(records[n : n+4])
			records[targetPosition] = inputValue1 + inputValue2
			// fmt.Println(records[n : n+4])
			// fmt.Println("mode1: ", mode1)
			// fmt.Println("mode2: ", mode2)
			// fmt.Println("input1:", inputValue1)
			// fmt.Println("input2:", inputValue2)
			// fmt.Println("Value: ", records[targetPosition])
			// fmt.Println("Target:", targetPosition)
			n += 4
		case 2:
			targetPosition := records[n+3]
			inputValue1 := calcInputValue(records, mode1, n+1)
			inputValue2 := calcInputValue(records, mode2, n+2)
			// fmt.Println(records[n : n+4])
			records[targetPosition] = inputValue1 * inputValue2
			// fmt.Println(records[n : n+4])
			// fmt.Println("mode1: ", mode1)
			// fmt.Println("mode2: ", mode2)
			// fmt.Println("input1:", inputValue1)
			// fmt.Println("input2:", inputValue2)
			// fmt.Println("Value: ", records[targetPosition])
			// fmt.Println("Target:", targetPosition)
			n += 4
		case 3:
			targetPosition := records[n+1]
			records[targetPosition] = input
			// fmt.Println(records[n : n+2])
			// fmt.Println("Value: ", records[targetPosition])
			// fmt.Println("Target:", targetPosition)
			n += 2
		case 4:
			inputValue1 := calcInputValue(records, mode1, n+1)
			// fmt.Println(records[n : n+2])
			fmt.Println("Output: ", inputValue1)
			// fmt.Println("Target:", targetPosition)
			n += 2
		}
		x := 0
		for x < 0 {
			opcode := records[x] % 100
			switch opcode {
			case 1:
				fmt.Println(x, records[x:x+4])
				x += 4
			case 2:
				fmt.Print(x, records[x:x+4])
				x += 4
			case 3:
				fmt.Print(x, records[x:x+2])
				x += 2
			case 4:
				fmt.Print(x, records[x:x+2])
				x += 2
			case 99:
				if x >= noRecords-5 {
					fmt.Print(x, records[x:])
					x = noRecords + 1
				} else {
					fmt.Print(x, records[x:x+4])
					x += 4
				}
			default:
				fmt.Print(x, records[x:x+4])
				x += 4
			}

		}
		// fmt.Println()
		// fmt.Println("------------------------------------------------------------")
	}
}

func calcInputValue(records []int, mode int, position int) (value int) {
	var inputPosition int
	switch mode {
	case 0:
		inputPosition = records[position]
	case 1:
		inputPosition = position
	}
	return records[inputPosition]
}

func stringValuesToInts(stringValues []string) ([]int, error) {
	var err error
	values := make([]int, len(stringValues))
	for i := range values {
		values[i], err =
			strconv.Atoi(stringValues[i])
		if err != nil {
			fmt.Println(err)
			return values, err
		}
	}
	return values, nil
}
