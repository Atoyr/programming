package main

import "fmt"

func q2() {
	nextReader = newScanner()
	n := nextInt()
	N := nextInts(n)
	count := 0
	for i := 1; i < n-1; i++ {
		if N[i-1] < N[i] && N[i+1] < N[i] {
			count++
		}
	}
	fmt.Println(count)
}
