package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	count := 0
	for i := 172851; i < 675869; i++ {
		if check(i) == false {
			continue
		}
		count++
	}
	fmt.Println(count)
	count2 := 0
	for i := 172851; i < 675869; i++ {
		if check2(i) == false {
			continue
		}
		count2++
	}
	fmt.Println(count2)
}

// solution for first question, a bit long..
func check(num int) bool {
	num2 := 0
	count := 0
	s := strconv.Itoa(num)
	for i, elem := range s {
		num1 := num2
		num2 = int(elem - '0')
		if i == 0 {
			continue
		}
		if num1 > num2 {
			return false
		}
		if num1 == num2 {
			count++
		}
	}
	if count == 0 {
		return false
	}
	return true
}

// solution for second question, a different approach for the fun of it
func check2(num int) bool {
	s := strconv.Itoa(num)
	var a []int
	for _, elem := range s {
		a = append(a, int(elem-'0'))
	}
	if sort.IntsAreSorted(a) == false {
		return false
	}

	match := false
	counter := make(map[int]int)
	for _, row := range a {
		counter[row]++
	}
	for _, v := range counter {
		// change check to 'v == 1' for a solution for question 1
		if v != 2 {
			continue
		}
		match = true
	}

	return match
}
