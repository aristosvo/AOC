package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode/utf8"
)

//Coordinate is used to save locations
type Coordinate struct {
	x int
	y int
}

func main() {
	var filename = flag.String("file", "../input.txt", "file to analyze")
	flag.Parse()

	csvfile, err := os.Open(*filename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()

	// Parse the file to two wires, called Red and Blue
	r := csv.NewReader(csvfile)
	routeRed, err := r.Read()
	coordinatesRed := routeToCoordinates(routeRed)

	routeBlue, err := r.Read()
	coordinatesBlue := routeToCoordinates(routeBlue)

	distanceToCrossing := 0

	for i := 1; i < len(coordinatesRed); i++ {
		for j := 1; j < len(coordinatesBlue); j++ {
			result, coordinate := crossing2d(coordinatesRed[i-1], coordinatesRed[i],
				coordinatesBlue[j-1], coordinatesBlue[j])
			if result == false {
				continue
			}
			manhattanDistance := abs(coordinate.x) + abs(coordinate.y)
			if distanceToCrossing < manhattanDistance && distanceToCrossing != 0 {
				continue
			}
			distanceToCrossing = manhattanDistance
		}
	}
	fmt.Println(distanceToCrossing)
}

// crossing2d checks if there is a crossing between two lines (2x2 coordinates) and returns the coordinate
func crossing2d(cRed1, cRed2, cBlue1, cBlue2 Coordinate) (bool, Coordinate) {
	crossX, coordinateX := cross1d(cRed1.x, cRed2.x, cBlue1.x, cBlue2.x)
	crossY, coordinateY := cross1d(cRed1.y, cRed2.y, cBlue1.y, cBlue2.y)
	if !crossY || !crossX {
		return false, Coordinate{x: 0, y: 0}
	}

	return true, Coordinate{x: coordinateX, y: coordinateY}
}

// cross1d checks if there is a crossing between two lines in one dimension and returns the value
func cross1d(vRed1, vRed2, vBlue1, vBlue2 int) (crossed bool, value int) {
	maxRed := max(vRed1, vRed2)
	minRed := min(vRed1, vRed2)
	maxBlue := max(vBlue1, vBlue2)
	minBlue := min(vBlue1, vBlue2)

	// Does Red enclose Blue or Blue enclose Red?
	if maxRed > maxBlue && minRed < minBlue {
		crossed = true
		value = maxBlue
	} else if maxBlue > maxRed && minBlue < minRed {
		crossed = true
		value = maxRed
	}

	return crossed, value
}

// routeToCoordinates makes it easier to calculate crossings by reducing the route to coordinates
func routeToCoordinates(route []string) []Coordinate {
	var coordinates []Coordinate
	coordinates = append(coordinates, Coordinate{x: 0, y: 0})
	for i := 0; i < len(route); i++ {
		direction, dist := directionAndDistance(route[i])
		x := coordinates[i].x
		y := coordinates[i].y
		switch direction {
		case 'R':
			coordinates = append(coordinates, Coordinate{x: x + dist, y: y})
		case 'L':
			coordinates = append(coordinates, Coordinate{x: x - dist, y: y})
		case 'U':
			coordinates = append(coordinates, Coordinate{x: x, y: y + dist})
		case 'D':
			coordinates = append(coordinates, Coordinate{x: x, y: y - dist})
		}
	}
	return coordinates
}

// directionAndDistance gives the direction and distance back
func directionAndDistance(s string) (rune, int) {
	direction := []rune(s)[0]
	_, i := utf8.DecodeRuneInString(s)
	dist, _ := strconv.Atoi(s[i:])

	return direction, dist
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
