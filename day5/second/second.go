package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var filename = flag.String("file", "../input.txt", "file to analyze")
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

	IntCodeComputer(records)
}

// IntCodeComputer yeah
func IntCodeComputer(records []int) {

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
		switch opcode {
		case 1:
			targetPosition := records[n+3]
			parameterValue1 := calcParamValue(records, mode1, n+1)
			parameterValue2 := calcParamValue(records, mode2, n+2)
			records[targetPosition] = parameterValue1 + parameterValue2
			n += 4
		case 2:
			targetPosition := records[n+3]
			parameterValue1 := calcParamValue(records, mode1, n+1)
			parameterValue2 := calcParamValue(records, mode2, n+2)
			records[targetPosition] = parameterValue1 * parameterValue2
			n += 4
		case 3:
			targetPosition := records[n+1]
			fmt.Println("Give input:")
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSuffix(input, "\n")
			inputt, err := strconv.Atoi(input)
			if err != nil {
				log.Fatalln("You did something stupid", err)
			}
			records[targetPosition] = inputt
			n += 2
		case 4:
			parameterValue1 := calcParamValue(records, mode1, n+1)
			fmt.Println("Output:", parameterValue1)
			n += 2
		case 5:
			parameterValue1 := calcParamValue(records, mode1, n+1)
			parameterValue2 := calcParamValue(records, mode2, n+2)
			if parameterValue1 != 0 {
				n = parameterValue2
			} else {
				n += 3
			}
		case 6:
			parameterValue1 := calcParamValue(records, mode1, n+1)
			parameterValue2 := calcParamValue(records, mode2, n+2)
			if parameterValue1 == 0 {
				n = parameterValue2
			} else {
				n += 3
			}
		case 7:
			targetPosition := records[n+3]
			parameterValue1 := calcParamValue(records, mode1, n+1)
			parameterValue2 := calcParamValue(records, mode2, n+2)
			if parameterValue1 < parameterValue2 {
				records[targetPosition] = 1
			} else {
				records[targetPosition] = 0
			}
			n += 4
		case 8:
			targetPosition := records[n+3]
			parameterValue1 := calcParamValue(records, mode1, n+1)
			parameterValue2 := calcParamValue(records, mode2, n+2)
			if parameterValue1 == parameterValue2 {
				records[targetPosition] = 1
			} else {
				records[targetPosition] = 0
			}
			n += 4
		}

	}
}

func calcParamValue(records []int, mode int, position int) (value int) {
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
