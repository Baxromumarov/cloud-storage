package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(len(arr), cap(arr)) // 6 6
	b := arr[1:2]
	a := arr[2:4]
	fmt.Println(len(b), cap(b)) // 1 5
	fmt.Println(len(a), cap(a)) // 2 4

}
