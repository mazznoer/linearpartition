package linearpartition

import "testing"

func TestFloat64(t *testing.T) {
	seq := []float64{0.75, 1.2, 0.11, 0.2, 0.9, 0.52, 0.33, 0.20, 0.57, 0.6, 0.17}

	var result [][]float64
	var expected [][]float64

	result = Float64(seq, 0)

	if len(result) != 0 {
		t.Errorf("It should 0")
	}
	t.Log(result)

	expected = [][]float64{{0.75}, {1.2}, {0.11}, {0.2}, {0.9}, {0.52}, {0.33}, {0.2}, {0.57}, {0.6}, {0.17}}
	result = Float64(seq, len(seq))

	if !isEqualFloat64(result, expected) {
		t.Errorf("It expected: %v, but got: %v", expected, result)
	}
	t.Log(result)

	result = Float64(seq, len(seq)+1)

	if !isEqualFloat64(result, expected) {
		t.Errorf("It expected: %v, but got: %v", expected, result)
	}
	t.Log(result)

	expected = [][]float64{
		{0.75, 1.2},
		{0.11, 0.2, 0.9, 0.52},
		{0.33, 0.2, 0.57, 0.6, 0.17},
	}
	result = Float64(seq, 3)

	if !isEqualFloat64(result, expected) {
		t.Errorf("It expected: %v, but got: %v", expected, result)
	}
	t.Log(result)
}

func TestInt(t *testing.T) {
	seq := []int{1, 6, 2, 3, 1, 4, 2, 2, 1, 3}

	var result [][]int
	var expected [][]int

	result = Int(seq, 0)

	if len(result) != 0 {
		t.Errorf("It should 0")
	}
	t.Log(result)

	expected = [][]int{{1}, {6}, {2}, {3}, {1}, {4}, {2}, {2}, {1}, {3}}
	result = Int(seq, len(seq))

	if !isEqualInt(result, expected) {
		t.Errorf("It expected: %v, but got: %v", expected, result)
	}
	t.Log(result)

	result = Int(seq, len(seq)+1)

	if !isEqualInt(result, expected) {
		t.Errorf("It expected: %v, but got: %v", expected, result)
	}
	t.Log(result)

	expected = [][]int{
		{1, 6, 2},
		{3, 1, 4},
		{2, 2, 1, 3},
	}
	result = Int(seq, 3)

	if !isEqualInt(result, expected) {
		t.Errorf("It expected: %v, but got: %v", expected, result)
	}
	t.Log(result)
}

func isEqualFloat64(a, b [][]float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, av := range a {
		if len(av) != len(b[i]) {
			return false
		}
		for j, v := range av {
			if v != b[i][j] {
				return false
			}
		}
	}
	return true
}

func isEqualInt(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, av := range a {
		if len(av) != len(b[i]) {
			return false
		}
		for j, v := range av {
			if v != b[i][j] {
				return false
			}
		}
	}
	return true
}
