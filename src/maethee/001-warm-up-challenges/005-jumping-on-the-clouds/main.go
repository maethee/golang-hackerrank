package main

import "fmt"

func main() {
	clounds := []int32{0, 0, 0, 1, 0, 0}
	routes := jumpingOnClouds(clounds)
	fmt.Printf("%v \n", routes)

}

func predictNextJump(dis int, c []int32) int {
	if dis+2 < len(c) && c[dis+2] == 0 {
		return 2
	}
	return 1
}

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {
	var jumps int
	var dis int
	for {
		if dis >= len(c)-1 { //-1 because we compare with index that start at 0
			break
		}
		dis += predictNextJump(dis, c)
		jumps += 1
		fmt.Printf("%v %v %v \n", jumps, dis, len(c))
	}

	return int32(jumps)
}
