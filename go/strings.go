package main

import "fmt"
import "strings"

func find(base string, search string) {
	// retrun true or false
	fmt.Println(strings.Contains(base, search))
}

func upperLower(str string) {
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
}

func trim(base string, trimstr string) {
	fmt.Println(strings.TrimLeft(base, trimstr))
	fmt.Println(strings.TrimRight(base, trimstr))
	fmt.Println(strings.Trim(base, trimstr))
}

func replace(base string, before string, after string) {
	fmt.Println(strings.Replace(base, before, after, -1))
}

func split(base string, splitstr string) {
	fmt.Println(strings.Split(base, splitstr))
}

func join(strs []string, sep string) {
	fmt.Println(strings.Join(strs, sep))
}
