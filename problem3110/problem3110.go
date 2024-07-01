package problem3110

import (
	"fmt"

	"github.com/lebrancconvas/Golang-Leetcode/problem3110/utils"
)

func Problem3110(s string) int {
	var asciiStore []int
	for i := 0; i < len(s); i++ {
		ascii := int(s[i])
		asciiStore = append(asciiStore, ascii)
	}

	fmt.Println(asciiStore)

	var score int = 0
	for i := 0; i < len(asciiStore)-1; i++ {
		score += utils.Abs(asciiStore[i] - asciiStore[i+1])
	}

	return score

}
