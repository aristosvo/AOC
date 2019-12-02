package main

import (
	"bufio"
	"log"
	"os"
	"fmt"
	"strconv"
	"flag"
)

func main() {

	var filename = flag.String("file", "input.txt", "file to analyze")

	flag.Parse()
    fmt.Println(*filename)
}


