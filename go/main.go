package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	nextReader = newScanner()
	N := nextInt()
	W := nextInt()
	values := make([]int, 110) // [0 0 0]
	weights := make([]int, 110) // [0 0 0]

	dp := make([][]int, 110) // [0 0 0]
	for i := range dp {
		dp[i] = make([]int, 10100) // [0 0 0]
	}

	
	for i :=0 ; i < N ; i++{
		values[i] = nextInt()
		weights[i] = nextInt()
	}

	for i := 0; i < N ; i++{
		for w := 0; w <= W ; w++{
			if (w >= weights[i]) {
				dp[i+1][w] = maxInt(dp[i][w-weights[i]] + values[i], dp[i][w])
			}else {
				dp[i+1][w] = dp[i][w]
			}
		}
	}
	fmt.Print(dp[N][W])
}

var nextReader func() string

func newScanner() func() string {
	r := bufio.NewScanner(os.Stdin)
	r.Buffer(make([]byte, 1024), int(1e+11))
	r.Split(bufio.ScanWords)
	return func() string {
		r.Scan()
		return r.Text()
	}
}
func nextString() string {
	return nextReader()
}
func nextInt64() int64 {
	v, _ := strconv.ParseInt(nextReader(), 10, 64)
	return v
}
func nextInt() int {
	v, _ := strconv.Atoi(nextReader())
	return v
}
func nextInts(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = nextInt()
	}
	return r
}
func nextInt64s(n int) []int64 {
	r := make([]int64, n)
	for i := 0; i < n; i++ {
		r[i] = nextInt64()
	}
	return r
}
func nextFloat64() float64 {
	f, _ := strconv.ParseFloat(nextReader(), 64)
	return f
}

func maxInt(a,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
