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
	max := 0
	for _, a := range permutations([]int{0, 1, 2, 3, 4}) {
		records, err := stringValuesToInts(inputRecords)
		if err != nil {
			log.Fatalln("You did something stupid", err)
		}
		input := 0
		for _, v := range a {
			input = IntCodeComputer(records, []int{v, input})
		}
		if input <= max {
			continue
		}
		max = input
	}
	fmt.Println(max)
}

// IntCodeComputer yeah
func IntCodeComputer(records []int, input []int) int {
	inputCount := 0
	n := 0
	for n < len(records) {
		opcode := records[n] % 100
		//probably not used for now
		mode3 := records[n] / 10000
		mode2 := (records[n] - mode3*10000) / 1000
		mode1 := (records[n] - mode2*1000) / 100

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
			records[targetPosition] = input[inputCount]
			inputCount++
			n += 2
		case 4:
			parameterValue1 := calcParamValue(records, mode1, n+1)
			return parameterValue1
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
		case 99:
			return 0
		}
	}
	return 0
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

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}
