package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	foo(arr)
	fmt.Println(arr)
}

func foo(arr []int) {
	arr[0] = 100
}
