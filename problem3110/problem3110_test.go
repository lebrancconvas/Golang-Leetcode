package problem3110

import "testing"

type TestCase struct {
	Name  string
	Input string
	Want  int
}

func TestProblem3110(t *testing.T) {
	testCases := []TestCase{
		{Name: "hello should return 13", Input: "hello", Want: 13},
		{Name: "zaz should return 50", Input: "zaz", Want: 50},
		{Name: "Emma should return 52", Input: "Emma", Want: 52},
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
