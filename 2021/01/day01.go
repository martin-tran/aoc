package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var inFile = flag.String("i", "", "Input file to read from.")

func readInputs() (depths []int) {
	input, err := os.Open(*inFile)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		depth, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, depth)
	}

	return depths
}

func solve(depths []int, winSize int) (ans int) {
	i := 0
	for ; i < len(depths)-winSize; i++ {
		if depths[i] < depths[i+winSize] {
			ans++
		}
	}

	return ans
}

func main() {
	flag.Parse()
	depths := readInputs()

	fmt.Println("Part 1:", solve(depths, 1))
	fmt.Println("Part 2:", solve(depths, 3))
}
