package problem3146

import "testing"

type TestCase struct {
	Name  string
	Input1 string
	Input2 string
	Want  int
}

func TestProblem3110(t *testing.T) {
	testCases := []TestCase{
		{Name: "abc and bca should return 2", Input1: "abc", Input2: "bca", Want: 2},
		{Name: "abcde and edbac should return 12", Input1: "abcde", Input2: "edbac", Want: 12},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			got := Problem3146(tt.Input1, tt.Input2)

			if got != tt.Want {
				t.Errorf("[ERROR] %v: got %v, but want %v", tt.Name, got, tt.Want)
			}
		})
	}
}
