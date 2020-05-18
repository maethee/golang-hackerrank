package main

import "fmt"

//https://www.hackerrank.com/challenges/sock-merchant/problem

func main() {

	steps := "UDDDUDUU"
	x := countingValleys(8, steps)
	fmt.Printf("%v \n", x)

}

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {
	var seaLevel int32
	var countingValleys int32
	step := []byte(s)
	for _, v := range step {
		if v == 85 {
			seaLevel += 1
			if seaLevel == 0 {
				countingValleys += 1
			}
		}
		if v == 68 {
			seaLevel += -1
		}
	}
	return countingValleys
}
