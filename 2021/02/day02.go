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

const (
	forward = iota
	down
	up
)

type cmd struct {
	dir, magnitude int
}

func readInputs() (cmds []cmd) {
	input, err := os.Open(*inFile)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		dir := scanner.Text()
		scanner.Scan()
		magnitude, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		switch dir[0] {
		case 'f':
			cmds = append(cmds, cmd{forward, magnitude})
		case 'd':
			cmds = append(cmds, cmd{down, magnitude})
		case 'u':
			cmds = append(cmds, cmd{up, magnitude})
		}
	}

	return cmds
}

func solve1(cmds []cmd) int {
	var pos, depth int
	for _, c := range cmds {
		switch c.dir {
		case forward:
			pos += c.magnitude
		case down:
			depth += c.magnitude
		case up:
			depth -= c.magnitude
		}
	}

	return pos * depth
}

func solve2(cmds []cmd) int {
	var pos, depth, aim int
	for _, c := range cmds {
		switch c.dir {
		case forward:
			pos += c.magnitude
			depth += aim * c.magnitude
		case down:
			aim += c.magnitude
		case up:
			aim -= c.magnitude
		}
	}

	return pos * depth
}

func main() {
	flag.Parse()

	cmds := readInputs()
	fmt.Println("Part 1:", solve1(cmds))
	fmt.Println("Part 2:", solve2(cmds))
}
