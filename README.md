# linearpartition

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/mazznoer/linearpartition?tab=doc)
[![Build Status](https://travis-ci.org/mazznoer/linearpartition.svg?branch=master)](https://travis-ci.org/mazznoer/linearpartition)
[![Build Status](https://github.com/mazznoer/linearpartition/workflows/Go/badge.svg)](https://github.com/mazznoer/linearpartition/actions)
[![go report](https://goreportcard.com/badge/github.com/mazznoer/linearpartition)](https://goreportcard.com/report/github.com/mazznoer/linearpartition)
[![codecov](https://codecov.io/gh/mazznoer/linearpartition/branch/master/graph/badge.svg)](https://codecov.io/gh/mazznoer/linearpartition)

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

Ported from [Javascript implementation](https://github.com/azrosen92/linear-partition-javascript).
