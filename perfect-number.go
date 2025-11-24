package main

import (
	"fmt"
	"log"
)

func main() {
	// fmt.Println(FindPerfectNuber(0, 000))
	// fmt.Println(isPerfectNumber(6))
	fmt.Println(predictComputationsForRange(1, 8200))
}

const (
	MAX_COMPUTATIONS_PER_THREAD = 3_000_000
)

// func findMostSuitable(start int, finish int) ([2]int, int) {
// out := [2]int{}
// delta := 0
// prevCmpt := 0
// cmpt := predictComputationsForRange(start, finish)
// if cmpt >=
// for cmpt < MAX_COMPUTATIONS_PER_THREAD {

// 	cmpt := predictComputationsForRange(start, finish-delta)
// 	prevCmpt = cmpt
// 	delta += ((MAX_COMPUTATIONS_PER_THREAD - cmpt) / 2) + 1

// }
// return put, prevCmpt
// }

func FindPerfectNuber(start int, finish int) ([]int, int) {
	out := make([]int, 1)
	computed := 0
	for i := start; i <= finish; i++ {
		isPefect, cmptd := isPerfectNumber(i)
		if isPefect {
			out = append(out, i)
		}
		computed += cmptd
	}
	return out, computed
}

func isPerfectNumber(n int) (bool, int) {
	computations := 0
	sum := 0
	for i := 1; i < getHalf(n); i++ {
		if n%i == 0 {
			sum += i
		}
		computations++
	}
	return sum == n, computations
}

func getHalf(n int) int {
	return (n / 2) + 1
}

func predictComputations(n int) int {
	return getHalf(n) - 1
}

func predictComputationsForRange(start int, finish int) int {
	predictedCmpt := 0
	for i := start; i <= finish; i++ {
		predictedCmpt = predictComputations(i)
		found, realCompt := isPerfectNumber(i)
		if found {
			fmt.Println("Found!!!", i)
		}
		if realCompt != predictedCmpt {
			log.Fatalf("WRONG!!! real computations:%d vs predicted:%d for %d\n", realCompt, predictedCmpt, i)
		}
	}
	return predictedCmpt
}
