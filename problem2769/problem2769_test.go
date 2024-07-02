package problem2769

import "testing"

type TestCase struct {
	Name   string
	Input1 int
	Input2 int
	Want   int
}

func TestProblem3110(t *testing.T) {
	testCases := []TestCase{
		{Name: "num = 4, t = 1 should return 6", Input1: 4, Input2: 1, Want: 6},
		{Name: "num = 3, t = 2 should return 7", Input1: 3, Input2: 2, Want: 7},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			got := Problem2769(tt.Input1, tt.Input2)

			if got != tt.Want {
				t.Errorf("[ERROR] %v: got %v, but want %v", tt.Name, got, tt.Want)
			}
		})
	}
}
