package tdma

func Solve(a, b, c, d []float64) []float64 {
	n := len(a)

	c[0] /= b[0]
	d[0] /= b[0]
	n--
	for i := 1; i < n; i++ {
		c[i] /= b[i] - a[i]*c[i-1]
		d[i] = (d[i] - a[i]*d[i-1]) / (b[i] - a[i]*c[i-1])
	}

	d[n] = (d[n] - a[n]*d[n-1]) / (b[n] - a[n]*c[n-1])

	for i := n - 1; i >= 0; i-- {
		d[i] -= c[i] * d[i+1]
	}

	return d
}
