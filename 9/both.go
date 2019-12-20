package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type cpu struct {
	mem     []int
	pos     int
	relBase int
}

func main() {
	var filename = flag.String("file", "input.txt", "file to analyze")
	flag.Parse()

	csvfile, err := os.Open(*filename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()
	basicOperationOfSystemTest, err := csv.NewReader(csvfile).Read()
	if err != nil {
		log.Fatalln("You did something stupid", err)
	}

	fmt.Println("Question 1:")
	fmt.Println(strings.Repeat("-", 30), "\nBOOST keycode of test run:")
	CPUMem1 := fillCPUMem(basicOperationOfSystemTest)
	IntCodeComputer(&CPUMem1, 1)

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("\nQuestion 2:", "\n"+strings.Repeat("-", 30), "\nThe coordinates of the distress signal:")
	CPUMem2 := fillCPUMem(basicOperationOfSystemTest)
	IntCodeComputer(&CPUMem2, 2)

}

func fillCPUMem(records []string) cpu {
	mem, err := stringValuesToInts(records)
	mem = append(mem, make([]int, 10000)...)
	if err != nil {
		log.Fatalln("You did something stupid", err)
	}
	return cpu{mem: mem}
}

// IntCodeComputer yeah
func IntCodeComputer(CPU *cpu, input int) (output int, finished bool) {

	for (*CPU).pos < len((*CPU).mem) {

		opcode := (*CPU).mem[(*CPU).pos] % 100
		mode3 := (*CPU).mem[(*CPU).pos] / 10000
		mode2 := ((*CPU).mem[(*CPU).pos] - mode3*10000) / 1000
		mode1 := ((*CPU).mem[(*CPU).pos] - mode2*1000 - mode3*10000) / 100

		switch opcode {
		case 1:
			targetPosition := calcTargetPosition((*CPU).mem, mode3, (*CPU).pos+3, &(*CPU).relBase)
			paramValue1 := calcParamValue((*CPU).mem, mode1, (*CPU).pos+1, &(*CPU).relBase)
			paramValue2 := calcParamValue((*CPU).mem, mode2, (*CPU).pos+2, &(*CPU).relBase)
			(*CPU).mem[targetPosition] = paramValue1 + paramValue2
			(*CPU).pos += 4
		case 2:
			targetPosition := calcTargetPosition((*CPU).mem, mode3, (*CPU).pos+3, &(*CPU).relBase)
			paramValue1 := calcParamValue((*CPU).mem, mode1, (*CPU).pos+1, &(*CPU).relBase)
			paramValue2 := calcParamValue((*CPU).mem, mode2, (*CPU).pos+2, &(*CPU).relBase)
			(*CPU).mem[targetPosition] = paramValue1 * paramValue2
			(*CPU).pos += 4
		case 3:
			targetPosition := calcTargetPosition((*CPU).mem, mode1, (*CPU).pos+1, &(*CPU).relBase)
			// if inputCount >= len(input) {
			// 	break CPURun
			// }
			(*CPU).mem[targetPosition] = input
			//inputCount++
			(*CPU).pos += 2
		case 4:
			paramValue1 := calcParamValue((*CPU).mem, mode1, (*CPU).pos+1, &(*CPU).relBase)
			//fmt.Println((*CPU).mem, mode1, paramValue1)
			fmt.Println("Output:", paramValue1)
			(*CPU).pos += 2
		case 5:
			paramValue1 := calcParamValue((*CPU).mem, mode1, (*CPU).pos+1, &(*CPU).relBase)
			paramValue2 := calcParamValue((*CPU).mem, mode2, (*CPU).pos+2, &(*CPU).relBase)
			if paramValue1 != 0 {
				(*CPU).pos = paramValue2
			} else {
				(*CPU).pos += 3
			}
		case 6:
			paramValue1 := calcParamValue((*CPU).mem, mode1, (*CPU).pos+1, &(*CPU).relBase)
			paramValue2 := calcParamValue((*CPU).mem, mode2, (*CPU).pos+2, &(*CPU).relBase)
			if paramValue1 == 0 {
				(*CPU).pos = paramValue2
			} else {
				(*CPU).pos += 3
			}
		case 7:

			paramValue1 := calcParamValue((*CPU).mem, mode1, (*CPU).pos+1, &(*CPU).relBase)
			paramValue2 := calcParamValue((*CPU).mem, mode2, (*CPU).pos+2, &(*CPU).relBase)
			targetPosition := calcTargetPosition((*CPU).mem, mode3, (*CPU).pos+3, &(*CPU).relBase)
			if paramValue1 < paramValue2 {
				(*CPU).mem[targetPosition] = 1
			} else {
				(*CPU).mem[targetPosition] = 0
			}
			(*CPU).pos += 4
		case 8:
			targetPosition := calcTargetPosition((*CPU).mem, mode3, (*CPU).pos+3, &(*CPU).relBase)
			paramValue1 := calcParamValue((*CPU).mem, mode1, (*CPU).pos+1, &(*CPU).relBase)
			paramValue2 := calcParamValue((*CPU).mem, mode2, (*CPU).pos+2, &(*CPU).relBase)
			if paramValue1 == paramValue2 {
				(*CPU).mem[targetPosition] = 1
			} else {
				(*CPU).mem[targetPosition] = 0
			}
			(*CPU).pos += 4
		case 9:
			paramValue1 := calcParamValue((*CPU).mem, mode1, (*CPU).pos+1, &(*CPU).relBase)
			(*CPU).relBase += paramValue1
			(*CPU).pos += 2
		case 99:
			return
		}
	}

	return
}

func calcParamValue(mem []int, mode, pos int, relBase *int) (value int) {
	var inputpos int

	switch mode {
	case 0:
		inputpos = mem[pos]
	case 1:
		inputpos = pos
	case 2:
		inputpos = mem[pos] + *relBase
	}

	return mem[inputpos]
}

func calcTargetPosition(mem []int, mode, pos int, relBase *int) int {
	var targetpos int

	switch mode {
	case 0:
		targetpos = mem[pos]
	case 2:
		targetpos = mem[pos] + *relBase
	}

	return targetpos
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
