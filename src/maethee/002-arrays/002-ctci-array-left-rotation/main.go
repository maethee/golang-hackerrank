package main

import "fmt"

/**
Reference: https://www.hackerrank.com/challenges/ctci-array-left-rotation
Idea: we can using mod (%) to find nex index to assign the value to new array
like:
total array = 5
assume number roration is 4
array value is [1,2,3,4,5]
the result [5,1,2,3,4]

value index
1 -> 1 (5-(4-index(0))%5 = 1
2 -> 2 (5-(4-index(1))%5 = 2
3 -> 3 (5-(4-index(2))%5 = 3
4 -> 4 (5-(4-index(3))%5 = 4
5 -> 0 (5-(4-index(4))%5 = 0


*/

func main() {

	arr := []int32{1, 2, 3, 4, 5}
	rotation := int32(4)
	ans := rotLeft(arr, rotation)
	fmt.Printf("Answer : %v", ans)
}

func getNewRotationIndex(roration int32, index int32, arrLen int32) int32 {
	return (arrLen - (roration - index)) % arrLen
}

// Complete the rotLeft function below.

func rotLeft(a []int32, d int32) []int32 {
	arrLen := int32(len(a))
	ans := make([]int32, arrLen)

	for index, v := range a {
		ni := getNewRotationIndex(d, int32(index), arrLen)
		ans[ni] = v
	}

	return ans
}
