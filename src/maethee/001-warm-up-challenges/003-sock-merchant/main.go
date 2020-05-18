package main

import "fmt"
//https://www.hackerrank.com/challenges/sock-merchant/problem

func main() {

	socks := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	x := sockMerchant(9, socks)
	fmt.Printf("%v \n", x)

}



// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {
	var ans int32
	pairs := map[int32]int32{}
	for _, v := range ar{
		pairs[v] += 1
	}

	for _, v := range pairs{
		ans += v/2
	}
	return ans
}