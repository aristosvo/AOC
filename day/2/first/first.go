package main

import (
	"flag"
	"fmt"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func main() {
	var filename = flag.String("file", "../input.txt", "file to analyze")
	flag.Parse()

	csvfile, err := os.Open(*filename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()
	input_records, err := csv.NewReader(csvfile).Read()
	if err != nil {
		log.Fatalln("You did something stupid", err)
	}
	
	// Starstruck preparations
	noun:= 12
	verb:= 2
	records, err:= stringValuesToInts(input_records)
	result:= IntCodeComputer(records, noun, verb) 
	fmt.Println("Result:", result)

}

// IntCodeComputer yeah
func IntCodeComputer(records []int, noun, verb int ) int {
	records[1] = noun
	records[2] = verb

	no_records := len(records)

	for i:=0; i+3 < no_records; i=4+i {
		opcode := records[i]
		if opcode == 99 {
			return records[0]
		}

		targetPosition := records[i+3]
		inputPosition1 := records[i+1]
		inputPosition2 := records[i+2]
		if inputPosition1 > no_records || inputPosition2 > no_records|| targetPosition > no_records {
			return 0
		}

		inputValue1 := records[inputPosition1]
		inputValue2 := records[inputPosition2]
		switch opcode {
		case 1:
			records[targetPosition] = inputValue1 + inputValue2
		case 2:
			records[targetPosition] = inputValue1 * inputValue2
		}
	}
	return records[0]
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
