package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func part1(in []int) int {
	mem := make([]int, len(in))
	copy(mem, in)
	mem[1] = 12
	mem[2] = 2

	return run(mem)
}

func part2(in []int) int {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			mem := make([]int, len(in))
			copy(mem, in)

			mem[1] = i
			mem[2] = j
			if k := run(mem); k == output {
				return 100*i + j
			}
		}
	}

	return -1
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Errorf("error reading input: %v\n", err)
		return
	}
	raw_input := scanner.Text()

	var in []int

	for _, n := range strings.Split(raw_input, ",") {
		i, err := strconv.Atoi(n)
		if err != nil {
			fmt.Errorf("error parsing input to int: %v\n", err)
			return
		}
		in = append(in, i)
	}

	fmt.Printf("Part 1: %v\nPart 2: %v\n", part1(in), part2(in))
}
