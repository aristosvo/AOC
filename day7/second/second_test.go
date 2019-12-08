package main

import (
	"testing"
)

func TestRunAmplifierCircuit(t *testing.T) {
	instructions := []int{9, 8, 7, 6, 5}
	amplifierControllerSoftware := []string{"3", "26", "1001", "26", "-4", "26", "3", "27", "1002", "27", "2", "27", "1", "27", "26", "27", "4", "27", "1001", "28", "-1", "28", "1005", "28", "6", "99", "0", "0", "5"}

	got := RunAmplifierCircuit(amplifierControllerSoftware, instructions)
	if got != 139629729 {
		t.Errorf("first position = %d, want %d", got, 139629729)
	}

}

func TestMaxAmplifierCircuit(t *testing.T) {
	instructionSet := []int{9, 8, 7, 6, 5}
	amplifierControllerSoftware := []string{"3", "26", "1001", "26", "-4", "26", "3", "27", "1002", "27", "2", "27", "1", "27", "26", "27", "4", "27", "1001", "28", "-1", "28", "1005", "28", "6", "99", "0", "0", "5"}

	got := MaxAmplifierCircuit(amplifierControllerSoftware, instructionSet)
	if got != 139629729 {
		t.Errorf("first position = %d, want %d", got, 139629729)
	}

}
