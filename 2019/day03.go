package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	key = map[string]Pair{
		"L": {-1, 0},
		"R": {1, 0},
		"U": {0, 1},
		"D": {0, -1},
	}
)

func abs(x int) int {
	if x > 0 {
		return x
	} else {
		return -x
	}
}

func min(arr []int) int {
	min := arr[0]
	for _, n := range arr {
		if n < min {
			min = n
		}
	}
	return min
}

type Instruction struct {
	Dir  string
	Dist int
}

func (instr Instruction) String() string {
	return fmt.Sprintf("%v%v", instr.Dir, instr.Dist)
}

type Pair struct {
	X, Y int
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.X, p.Y)
}

func part1(in1, in2 []Instruction, out chan int) {
	pos := Pair{0, 0}
	grid := map[Pair]bool{}
	for _, instr := range in1 {
		move := key[instr.Dir]
		for i := 0; i < instr.Dist; i++ {
			pos.X += move.X
			pos.Y += move.Y
			grid[Pair{pos.X, pos.Y}] = true
		}
	}

	pos = Pair{0, 0}
	var hits []int
	for _, instr := range in2 {
		move := key[instr.Dir]
		for i := 0; i < instr.Dist; i++ {
			pos.X += move.X
			pos.Y += move.Y
			if _, ok := grid[Pair{pos.X, pos.Y}]; ok {
				hits = append(hits, abs(pos.X)+abs(pos.Y))
			}
		}
	}

	out <- min(hits)
}

func part2(in1, in2 []Instruction, out chan int) {
	pos := Pair{0, 0}
	grid := map[Pair]int{}
	steps := 0
	for _, instr := range in1 {
		move := key[instr.Dir]
		for i := 0; i < instr.Dist; i++ {
			steps++
			pos.X += move.X
			pos.Y += move.Y
			grid[Pair{pos.X, pos.Y}] = steps
		}
	}

	pos = Pair{0, 0}
	steps = 0
	var hits []int
	for _, instr := range in2 {
		move := key[instr.Dir]
		for i := 0; i < instr.Dist; i++ {
			steps++
			pos.X += move.X
			pos.Y += move.Y
			if elem, ok := grid[Pair{pos.X, pos.Y}]; ok {
				hits = append(hits, steps+elem)
			}
		}
	}

	out <- min(hits)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	lineIdx := 0
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
			if data[i] == '\n' {
				lineIdx--
				return i + 1, nil, nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)
	
	lines := make([][]Instruction, 2)
	for scanner.Scan() {
		s := scanner.Text()
		if err := scanner.Err(); err != nil {
			fmt.Errorf("error reading input: %v\n", err)
			return
		}

		if lineIdx == -1 {
			lineIdx = 1
			continue
		}

		n, err := strconv.Atoi(s[1:])
		if err != nil {
			fmt.Errorf("error parsing input to int: %v\n", err)
			return
		}
		p := Instruction{string(s[0]), n}
		lines[lineIdx] = append(lines[lineIdx], p)
	}

	out1, out2 := make(chan int), make(chan int)
	go part1(lines[0], lines[1], out1)
	go part2(lines[0], lines[1], out2)

	fmt.Printf("Part 1: %v\nPart 2: %v\n", <-out1, <-out2)
}
