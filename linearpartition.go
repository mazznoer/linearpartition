package linearpartition

func Float64(seq []float64, k int) [][]float64 {
	n := len(seq)

	if k <= 0 {
		return [][]float64{}
	}

	if k >= n {
		res := make([][]float64, n)
		for i, v := range seq {
			res[i] = []float64{v}
		}
		return res
	}

	table := make([][]float64, n)
	solution := make([][]float64, n-1)

	for i := 0; i < n; i++ {
		table[i] = make([]float64, k)
		for j := 0; j < k; j++ {
			table[i][j] = 0
		}
	}

	for i := 0; i < n-1; i++ {
		solution[i] = make([]float64, k-1)
		for j := 0; j < k-1; j++ {
			solution[i][j] = 0
		}
	}

	for i := 0; i < n; i++ {
		if i != 0 {
			table[i][0] = seq[i] + table[i-1][0]
		} else {
			table[i][0] = seq[i]
		}
	}

	for j := 0; j < k; j++ {
		table[0][j] = seq[0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < k; j++ {
			var list_of_pairs [][]float64
			for x := 0; x < i; x++ {
				list_of_pairs = make([][]float64, i)
				for x := 0; x < i; x++ {
					max := fmax(table[x][j-1], table[i][0]-table[x][0])
					list_of_pairs[x] = []float64{max, float64(x)}
				}
			}

			m := list_of_pairs[0]
			for _, v := range list_of_pairs {
				if v[0] < m[0] {
					m = v
				}
			}

			table[i][j] = m[0]
			solution[i-1][j-1] = m[1]
		}
	}

	n -= 1
	k -= 2
	ans := [][]float64{}

	for {
		if k < 0 || n <= 0 { // xxx
			break
		}
		sub_list := []float64{}
		for i := int(solution[n-1][k]) + 1; i < n+1; i++ {
			sub_list = append(sub_list, seq[i])
		}
		ans = append([][]float64{sub_list}, ans...)
		n = int(solution[n-1][k])
		k--
	}

	beginning_list := make([]float64, n+1)
	for i := 0; i < n+1; i++ {
		beginning_list[i] = seq[i]
	}
	ans = append([][]float64{beginning_list}, ans...)
	return ans
}

func Int(seq []int, k int) [][]int {
	fseq := make([]float64, len(seq))
	for i, v := range seq {
		fseq[i] = float64(v)
	}
	fres := Float64(fseq, k)
	res := make([][]int, len(fres))
	for i, v := range fres {
		res[i] = make([]int, len(v))
		for j, x := range v {
			res[i][j] = int(x)
		}
	}
	return res
}

func fmax(a, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}
