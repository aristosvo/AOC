package main

import (
	"testing"
)

type Intcode struct {
	in []int
	answer int
}

var testinput  = []Intcode{
	{
		in: []int{1,9,10,3,2,3,11,0,99,30,40,50}, 
		answer: 3500,
	},
	{
		in: []int{1,0,0,0,99}, 
		answer: 2,
	},
	{
		in: []int{2,3,0,3,99}, 
		answer: 2,
	},
	{
		in: []int{1,1,1,4,99,5,6,0,99}, 
		answer: 30,
	},
}

func TestIntCodeComputer(t *testing.T) {
	for _, input := range testinput{

	noun := input.in[1]
	verb := input.in[2]
    got := IntCodeComputer(input.in, noun,verb )
    if got != input.answer {
        t.Errorf("first position = %d, want %d", got, input.answer)
	}
}
}