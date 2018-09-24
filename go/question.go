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

func q3() {
	nextReader = newScanner()
	nextInt()
	S := nextString()
	count := 0
	buf := ' '
	for _, c := range S {
		if c != buf {
			count++
			buf = c
		}
	}
	fmt.Println(count)
}

func q4() {
	nextReader = newScanner()
	nextInt()
	S := nextString()
	var count int64
	m := map[int32]int64{}
	for _, c := range S {
		_, ok := m[c]
		if ok {
			m[c]++
		} else {
			m[c] = 1
		}
	}
	for k, v := range m {
		v2, ok := m[k+1]
		if ok {
			count = count + v*v2
		}
	}
	fmt.Println(count)
}
