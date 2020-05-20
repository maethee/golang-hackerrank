package main

import "fmt"

/**
Reference: https://www.hackerrank.com/challenges/2d-array/problem
*/

func main() {
	arr := [][]int32{
		{-9, -9, -9, 1, 1, 1},
		{0, -9, 0, 4, 3, 2},
		{-9, -9, -9, 1, 2, 3},
		{0, 0, 8, 6, 6, 0},
		{0, 0, 0, -2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}

	x := hourglassSum(arr)
	fmt.Println("Ans %T : %v", x, x)
}

func sum(nums ...int32) int32 {
	var sum int32
	for _, v := range nums {
		sum += v
	}
	return sum
}

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	var ans int32 = 0
	var hgShape int = 3
	var x, y int = len(arr) - hgShape, len(arr[0]) - hgShape

	for i := 0; i <= x; i++ {
		for j := 0; j <= y; j++ {
			fr := sum(arr[i][j : j+hgShape]...)
			sr := arr[i+1][j+1]
			tr := sum(arr[i+2][j : j+hgShape]...)
			total := sum(fr, sr, tr)
			//I don,t want to store each sum of hourglass. just want to find high sum hourglass
			if total > ans || (j == 0 && i == 0) {
				ans = total
			}
		}
	}

	return ans
}
