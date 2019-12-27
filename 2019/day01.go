package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var (
	zero  = big.NewInt(int64(0))
	two   = big.NewInt(int64(2))
	three = big.NewInt(int64(3))
)

func part1(in chan big.Int, out chan string) {
	fuel := big.NewInt(int64(0))
	for mass := range in {
		mass.Div(&mass, three)
		mass.Sub(&mass, two)
		fuel.Add(fuel, &mass)
	}

	out <- fuel.String()

}

func part2(in chan big.Int, out chan string) {
	fuel := big.NewInt(int64(0))
	for mass := range in {
		for {
			mass.Sub(mass.Div(&mass, three), two)
			if mass.Cmp(zero) <= 0 {
				break
			}
			fuel.Add(fuel, &mass)
		}
	}

	out <- fuel.String()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	in1, in2 := make(chan big.Int), make(chan big.Int)
	out1, out2 := make(chan string), make(chan string)

	go part1(in1, out1)
	go part2(in2, out2)

	for {
		var mass1, mass2 big.Int
		if err := scanner.Scan(); err == false {
			close(in1)
			close(in2)
			break
		} else if err := scanner.Err(); err != nil {
			fmt.Errorf("error reading input: %v\n", err)
			return
		}
		mass1.SetString(scanner.Text(), 10)
		mass2.SetString(scanner.Text(), 10)
		
		in1 <- mass1
		in2 <- mass2
	}

	ans1 := <-out1
	ans2 := <-out2
	fmt.Printf("Part 1: %v\nPart 2: %v\n", ans1, ans2)
}
