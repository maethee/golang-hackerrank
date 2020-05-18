package main

import "fmt"

func main() {

	wording := "aba"
	numCha := int64(10)
	count := repeatedString(wording, numCha)
	fmt.Printf("%v \n", count)
}

func countA(s string) int64 {
	var count int64
	bs := []byte(s)
	for _, v := range bs {
		if v == 97 {
			count += 1
		}
	}
	return count
}

// Complete the repeatedString function below.
// Notes we don't create a real infinitly string and count a because we can make a logic to estimate count a that enough and reduce process and CPU process
// we don't do like sliceStr = append(sliceStr, s...) every repeate
func repeatedString(s string, n int64) int64 {
	var ans int64
	ls := int64(len(s))
	ca := countA(s)
	r := n/ls //repest
	rr := n-(r*ls) //repest rest
	sr := []byte(s)[:rr] //string rest

	ans = (r*ca) + countA(string(sr)) //result of (amount full repeat * count A) + CountA of rest string

	return ans
}
