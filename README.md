# linearpartition

Linear Partition in Go

```go
import "github.com/mazznoer/linearpartition"
```

```go
seq := []float64{1, 6, 2, 3, 1, 4, 2, 2, 1, 3}
result := linearpartition.Float64(seq, 3)

fmt.Println(result)
// Output: [[1 6 2] [3 1 4] [2 2 1 3]]
```

Ported from [Javascript](https://github.com/azrosen92/linear-partition-javascript)
