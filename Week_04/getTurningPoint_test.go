package week04

import (
	"fmt"
	"testing"
)

func TestGetTurningPoint(t *testing.T) {
	case1 := []int{6, 7, 0, 1, 2, 3, 4, 5}
	// case2 := []int{1, 2, 3, 4}
	res1 := getTurningPoint(case1)
	// res2 := getTurningPoint(case2)

	fmt.Println("case1:", case1, "res1:", res1)
	// fmt.Println("case2:", case2, "res2:", res2)
	// 	go test -v -run TestGetTurningPoint
	// === RUN   TestGetTurningPoint
	// case1: [4 5 6 7 0 1 2] res1: 4
	// case2: [1 2 3 4] res2: 0
	// --- PASS: TestGetTurningPoint (0.00s)
	// PASS
	// ok      github.com/TylerTang06/-algorithm015/Week_04    0.005s
}
