package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	foo(arr)
	fmt.Println(arr)

	arr2 := []int{1, 2, 3, 4, 5}
	foo2(arr2)
	fmt.Println(arr2)
}

func foo(arr []int) {
	arr[0] = 100
}

func foo2(arr2 []int) {
	arr2[0] = 100
}
