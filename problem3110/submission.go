package problem3110

func scoreOfString(s string) int {
  var asciiStore []int
	for i := 0; i < len(s); i++ {
		ascii := int(s[i])
		asciiStore = append(asciiStore, ascii)
	}

	var score int = 0
	for i := 0; i < len(asciiStore)-1; i++ {
		score += abs(asciiStore[i] - asciiStore[i+1])
	}

	return score
}


func abs(n int) int {
    if n < 0 {
        return n * -1
    }

    return n
}
