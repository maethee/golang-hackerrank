package main

/**
Reference: https://www.hackerrank.com/challenges/2d-array/problem
Idea:


*/
import "fmt"

func main() {

}

func sum(nums ...interface{}) int32 {
	var sum int32
	for _, v := range nums {
		sum += v
	}
	return sum
}

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {
	var hgShape int32 = 3
	var x, y int32 = len(arr) - hgShape, len(arr[0]) - hgShape

	//
}
