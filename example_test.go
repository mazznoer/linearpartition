package linearpartition_test

import "fmt"

func Example() {
	seq := []float64{1, 6, 2, 3, 1, 4, 2, 2, 1, 3}
	fmt.Println(linearpartition.Float64(seq, 3))
	// Output: [[1 6 2] [3 1 4] [2 2 1 3]]
}
