package main

import "fmt"

//reference : https://www.hackerrank.com/challenges/grading/problem
func main() {

	grades := []int32{4, 73, 67, 38, 33}

	r := gradingStudents(grades)

	fmt.Printf("Answer : %v \n", r)

}

func getGrade(score int32) int32 {
	round := -(score % 5) + 5
	fScore := score + round

	if round < 3 && fScore >= 40 {
		return fScore
	}

	return score
}

func gradingStudents(grades []int32) []int32 {
	var ans []int32
	for _, value := range grades {
		ans = append(ans, getGrade(value))
	}

	return ans
}
