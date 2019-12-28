package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	output = 19690720
)

func run(in []int) int {
	for i := 0; ; i += 4 {
		n := in[i]

		if n == 1 {
			in[in[i+3]] = in[in[i+1]] + in[in[i+2]]
		} else if n == 2 {
			in[in[i+3]] = in[in[i+1]] * in[in[i+2]]
		} else if n == 99 {
			return in[0]
		} else {
			return -1
		}
	}
}

func part1(in []int, out chan int) {
	mem := make([]int, len(in))
	copy(mem, in)
	mem[1] = 12
	mem[2] = 2

	out <- run(mem)
}

func part2(in []int, out chan int) {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			mem := make([]int, len(in))
			copy(mem, in)

			mem[1] = i
			mem[2] = j
			if k := run(mem); k == output {
				out <- 100*i + j
			}
		}
	}

	out <- -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)

	var in []int

	for scanner.Scan() {
		s := scanner.Text()

		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Errorf("error parsing input to int64: %v\n", err)
			return
		}
		in = append(in, i)
	}

	out1, out2 := make(chan int), make(chan int)
	go part1(in, out1)
	go part2(in, out2)

	fmt.Printf("Part 1: %v\nPart 2: %v\n", <-out1, <-out2)
}
