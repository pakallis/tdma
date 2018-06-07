# TDMA

A Go package that solves a tridiagonal system of equations using the Tridiagonal matrix algorithm (TDMA)

## Quick start
---
### Installation

```
go get -u github.com/pakallis/tdma/tdma
```
### Example Usage

1. Create file solve_tdma.go

```go
package main

import (
	"fmt"
	"github.com/pakallis/tdma/tdma"
)

func main() {
	var l = []float64{0, -1, -1, -1}
	var d = []float64{4, 4, 4, 4}
	var u = []float64{-1, -1, -1, 0}
	var b = []float64{5, 5, 10, 23}
	result := tdma.Solve(l, d, u, b)

	fmt.Println("Result:")
	for i := 0; i < len(result); i++ {
		fmt.Printf("%g ", result[i])
	}
	fmt.Println()
}

```

2. Run file with
```
go run solve_tdma.go
```

### Testing

```
cd tdma_test
go test
```
