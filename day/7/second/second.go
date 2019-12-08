package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type amplifier struct {
	mem []int
	pos int
}

func main() {
	var filename = flag.String("file", "../input.txt", "file to analyze")
	flag.Parse()

	csvfile, err := os.Open(*filename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()
	amplifierControllerSoftware, err := csv.NewReader(csvfile).Read()
	if err != nil {
		log.Fatalln("You did something stupid", err)
	}

	instructionSetOptions := []int{5, 6, 7, 8, 9}
	fmt.Println(MaxAmplifierCircuit(amplifierControllerSoftware, instructionSetOptions))
}

//MaxAmplifierCircuit finds the maximum output for the combination of software and instructionset options
func MaxAmplifierCircuit(amplifierControllerSoftware []string, instructionSetOptions []int) int {
	max := 0
	for _, instructionSet := range permutations(instructionSetOptions) {
		output := RunAmplifierCircuit(amplifierControllerSoftware, instructionSet)
		if output <= max {
			continue
		}
		max = output
	}
	return max
}

//RunAmplifierCircuit runs a circuit of amplifiers
func RunAmplifierCircuit(program []string, instruction []int) int {
	var ampsMem []amplifier
	for range instruction {
		ampsMem = append(ampsMem, fillAmpMem(program))
	}

	output := 0
	ending := false
	loop := 0
	var input []int
	for ending == false {
		for k, v := range instruction {
			if loop == 0 {
				input = []int{v, output}
			} else {
				input = []int{output}
			}
			output, ending = IntCodeComputer(&ampsMem[k], input)
		}
		loop++
	}
	return output
}

func fillAmpMem(records []string) amplifier {
	mem, err := stringValuesToInts(records)
	if err != nil {
		log.Fatalln("You did something stupid", err)
	}
	return amplifier{mem: mem}
}

// IntCodeComputer yeah
func IntCodeComputer(amp *amplifier, input []int) (output int, finished bool) {
	inputCount := 0

ampRun:
	for (*amp).pos < len((*amp).mem) {

		opcode := (*amp).mem[(*amp).pos] % 100
		mode3 := (*amp).mem[(*amp).pos] / 10000
		mode2 := ((*amp).mem[(*amp).pos] - mode3*10000) / 1000
		mode1 := ((*amp).mem[(*amp).pos] - mode2*1000) / 100

		switch opcode {
		case 1:
			targetPosition := (*amp).mem[(*amp).pos+3]
			paramValue1 := calcParamValue((*amp).mem, mode1, (*amp).pos+1)
			paramValue2 := calcParamValue((*amp).mem, mode2, (*amp).pos+2)
			(*amp).mem[targetPosition] = paramValue1 + paramValue2
			(*amp).pos += 4
		case 2:
			targetPosition := (*amp).mem[(*amp).pos+3]
			paramValue1 := calcParamValue((*amp).mem, mode1, (*amp).pos+1)
			paramValue2 := calcParamValue((*amp).mem, mode2, (*amp).pos+2)
			(*amp).mem[targetPosition] = paramValue1 * paramValue2
			(*amp).pos += 4
		case 3:
			targetPosition := (*amp).mem[(*amp).pos+1]
			if inputCount >= len(input) {
				break ampRun
			}
			(*amp).mem[targetPosition] = input[inputCount]
			inputCount++
			(*amp).pos += 2
		case 4:
			paramValue1 := calcParamValue((*amp).mem, mode1, (*amp).pos+1)
			output = paramValue1
			(*amp).pos += 2
		case 5:
			paramValue1 := calcParamValue((*amp).mem, mode1, (*amp).pos+1)
			paramValue2 := calcParamValue((*amp).mem, mode2, (*amp).pos+2)
			if paramValue1 != 0 {
				(*amp).pos = paramValue2
			} else {
				(*amp).pos += 3
			}
		case 6:
			paramValue1 := calcParamValue((*amp).mem, mode1, (*amp).pos+1)
			paramValue2 := calcParamValue((*amp).mem, mode2, (*amp).pos+2)
			if paramValue1 == 0 {
				(*amp).pos = paramValue2
			} else {
				(*amp).pos += 3
			}
		case 7:
			targetPosition := (*amp).mem[(*amp).pos+3]
			paramValue1 := calcParamValue((*amp).mem, mode1, (*amp).pos+1)
			paramValue2 := calcParamValue((*amp).mem, mode2, (*amp).pos+2)
			if paramValue1 < paramValue2 {
				(*amp).mem[targetPosition] = 1
			} else {
				(*amp).mem[targetPosition] = 0
			}
			(*amp).pos += 4
		case 8:
			targetPosition := (*amp).mem[(*amp).pos+3]
			paramValue1 := calcParamValue((*amp).mem, mode1, (*amp).pos+1)
			paramValue2 := calcParamValue((*amp).mem, mode2, (*amp).pos+2)
			if paramValue1 == paramValue2 {
				(*amp).mem[targetPosition] = 1
			} else {
				(*amp).mem[targetPosition] = 0
			}
			(*amp).pos += 4
		case 99:
			finished = true
			break ampRun
		}
	}

	return output, finished
}

func calcParamValue(mem []int, mode int, pos int) (value int) {
	var inputpos int
	switch mode {
	case 0:
		inputpos = mem[pos]
	case 1:
		inputpos = pos
	}
	return mem[inputpos]
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
