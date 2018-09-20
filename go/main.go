package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := loadStr()
	fmt.Println(s)
}

func loadStr() string {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}

func loadints(n int) []int {
	xs := make([]int, n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		sc.Scan()
		xs[i], _ = strconv.Atoi(sc.Text())
	}
	return xs
}

func loadLongStr(length int) string {
	len := length
	if len < 0 {
		len = 1000000
	}
	var rdr = bufio.NewReaderSize(os.Stdin, len)
	buf := make([]byte, 0, len)
	for {
		l, p, _ := rdr.ReadLine()
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}
