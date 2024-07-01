package problem3110

import "testing"

type TestCase struct {
	Name string
	Input string
	Want int
}

func TestProblem3110(t *testing.T) {
	testCases := []TestCase{
		{"hello should return 13", "hello", 13},
		{"zaz should return 50", "zaz", 50},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			got := Problem3110(tt.Input)

			if got != tt.Want {
				t.Errorf("[ERROR] %v: got %v, but want %v", tt.Name, got, tt.Want)
			}
		})
	}
}
