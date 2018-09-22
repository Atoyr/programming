package main

import "fmt"

func appendStr() {
	n := 10
	var m2 = make([]byte, 0, n)
	for i := 0; i < n; i++ {
		m2 = append(m2, nextString()...)
	}
	fmt.Println(string(m2))
}
