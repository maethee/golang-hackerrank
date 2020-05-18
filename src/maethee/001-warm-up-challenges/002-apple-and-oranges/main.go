package main

import "fmt"

//reference : https://www.hackerrank.com/challenges/apple-and-orange/problem
func main() {

	var s int32 = 2
	var t int32 = 3
	var a int32 = 1
	var b int32 = 5
	apples := []int32{2}
	oranges := []int32{-2}

	countApplesAndOranges(s,t,a,b,apples,oranges)
	
}


func isFallOnHouse (start int32, end int32, fall int32) bool {
	
		fmt.Printf("%v\n", fall)
	if(start <= fall && end >= fall){
		return true
	}
	return false
}
// Complete the countApplesAndOranges function below.
func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {

	var appleFOV int32
	var orangesFOV int32

	for _, v := range apples{
		fall := v+a
		if(isFallOnHouse(s,t,fall)){
			appleFOV+=1
		}
	}

	for _, v := range oranges{
		fall := v+b
		if(isFallOnHouse(s,t,fall)){
			orangesFOV+=1
		}
	}

	fmt.Printf("%v\n", appleFOV)
	fmt.Printf("%v\n", orangesFOV)

}
