package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

//Coordinate contains the x and y coordinate of a steroid
type Coordinate struct {
	x int
	y int
}

func main() {
	var filename = flag.String("file", "input.txt", "file to analyze")
	flag.Parse()

	file, err := os.Open(*filename)

	if err != nil {
		log.Fatalln("Couldn't open the file", err)
	}
	defer file.Close()

	m := createMapWithCoordinates(file)
	fmt.Println("Question 1:")
	fmt.Println(strings.Repeat("-", 30), "\nNumber of visible steroids:")
	max, steroid := findSteroidWithMaxView(m)
	fmt.Println(max)

	fmt.Println(strings.Repeat("-", 30))
	fmt.Println("\nQuestion 2:", "\n"+strings.Repeat("-", 30), "\nCode based on location of 200th destroyed steroid:")

	steroid200 := findNo200(m, steroid)
	fmt.Println(steroid200.x*100 + steroid200.y)
	fmt.Println(strings.Repeat("-", 30))
}

func findNo200(m map[Coordinate]bool, steroid Coordinate) Coordinate {
	_, view := findSteroidView(steroid, m)

	var ne []Coordinate
	var se []Coordinate
	var sw []Coordinate
	var nw []Coordinate
	for _, coordinate := range view {
		if coordinate.x >= steroid.x && coordinate.y <= steroid.y {
			ne = append(ne, coordinate)
		}
		if coordinate.x >= steroid.x && coordinate.y > steroid.y {
			se = append(se, coordinate)
		}
		if coordinate.x < steroid.x && coordinate.y > steroid.y {
			sw = append(sw, coordinate)
		}
		if coordinate.x < steroid.x && coordinate.y <= steroid.y {
			nw = append(nw, coordinate)
		}
	}

	sort.Slice(nw, func(i, j int) bool {
		return math.Atan(float64(steroid.y-nw[i].y)/float64(steroid.x-nw[i].x)) < math.Atan(float64(steroid.y-nw[j].y)/float64(steroid.x-nw[j].x))
	})

	steroid200 := nw[200-len(ne)-len(se)-len(sw)-1]
	return steroid200
}

func findSteroidWithMaxView(m map[Coordinate]bool) (max int, steroid Coordinate) {
	for coordinate := range m {
		count, _ := findSteroidView(coordinate, m)
		if count >= max {
			max = count
			steroid = coordinate
		}
	}
	return max, steroid
}

func findSteroidView(coordinate Coordinate, m map[Coordinate]bool) (int, []Coordinate) {
	var view []Coordinate
	for coordinateChecked := range m {
	S:
		switch {
		//Same coordinates shouldn't be counted
		case coordinateChecked.y == coordinate.y && coordinateChecked.x == coordinate.x:
			continue

		//Same y coordinate
		case coordinateChecked.y == coordinate.y:
			for _, x := range rangeInBetween(coordinateChecked.x, coordinate.x) {
				_, exists := m[Coordinate{x: x, y: coordinate.y}]
				if exists {
					break S
				}
			}
			view = append(view, coordinateChecked)

		//Same X coordinate
		case coordinateChecked.x == coordinate.x:
			for _, y := range rangeInBetween(coordinateChecked.y, coordinate.y) {
				_, exists := m[Coordinate{x: coordinate.x, y: y}]
				if exists {
					break S
				}
			}
			view = append(view, coordinateChecked)

		//Other
		default:
			for _, c := range findPossibleBlockersInBetween(coordinate, coordinateChecked) {
				_, exists := m[c]
				if exists {
					break S
				}
			}
			view = append(view, coordinateChecked)
		}
	}
	return len(view), view
}

func createMapWithCoordinates(file *os.File) map[Coordinate]bool {
	scanner := bufio.NewScanner(file)
	m := make(map[Coordinate]bool)
	count := 0
	for scanner.Scan() {
		for pos, char := range scanner.Text() {
			if char != '#' {
				continue
			}
			m[Coordinate{x: pos, y: count}] = true
		}
		count++
	}
	return m
}

func rangeInBetween(in1, in2 int) []int {
	min, max := minMax(in1, in2)
	a := make([]int, max-min-1)
	for i := range a {
		a[i] = min + i + 1
	}
	return a
}

func findPossibleBlockersInBetween(in1, in2 Coordinate) []Coordinate {
	var c []Coordinate
	diffX := in1.x - in2.x
	diffY := in1.y - in2.y
	absX := abs(diffX)
	absY := abs(diffY)
	min, _ := minMax(absX, absY)
	for _, divisionFactor := range rangeInBetween(1, min+1) {
		if absX%divisionFactor != 0 || absY%divisionFactor != 0 {
			continue
		}
		for _, factor := range rangeInBetween(0, divisionFactor) {
			coordinate := Coordinate{x: in1.x - (diffX/divisionFactor)*factor, y: in1.y - (diffY/divisionFactor)*factor}
			c = append(c, coordinate)
		}
	}
	return c
}

func minMax(x, y int) (int, int) {
	if x > y {
		return y, x
	}
	return x, y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
