package main

import (
	"fmt"
	"server/solver"
	"server/utils"
)

func main() {
	solvable, solution := solver.SolveBoard(65536)
	if solvable {
		fmt.Println("The puzzle can be solved by clicking on the following tiles:")
		for i := uint8(0); i < 25; i++ {
			if utils.TestBit(solution, i) {
				fmt.Printf("(%v, %v) ", i/5+1, i%5+1)
			}
		}
		fmt.Println()
	} else {
		fmt.Println("There is no solution")
	}
}
