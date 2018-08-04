package main

import (
	"fmt"
)

const (
	x  = iota
	y  = 1
	z  = iota
	pi = 3.1415
)

var s string = "hello"

func main() {
	fmt.Println(x, y, z, pi)

	fmt.Printf("%s\n", s)

	ss := "helloo"

	fmt.Printf("%s\n", ss)

	// 一维array
	a := [3]int{1, 2, 3}
	b := [10]int{1, 2, 3}
	c := [...]int{4, 5, 6}
	fmt.Print(a, b, c)

	// 二维array
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	fmt.Print(doubleArray, easyArray)

	// slice
	var fslice []int
	slice := []byte{'a', 'b', 'c'}

	fmt.Print(fslice, slice)

	// map
	var numbers map[string]int

	numbers = make(map[string]int)
	numbers["one"] = 1
	numbers["ten"] = 10
	numbers["three"] = 3

	fmt.Printf("第三个%d", numbers["three"])

	rating := map[string]float32{"C": 5, "Go": 4.5}

	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Salut" // 现在m["hello"]的值已经是Salut了

}
