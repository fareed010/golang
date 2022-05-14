package main

import (
	"fmt"
)

func main() {
	myArr := [6]int{1, 2, 3, 4, 5, 6}
	mySlice := myArr[2:6]
	
	fmt.Printf("mySlice = %v\n", mySlice)
	fmt.Printf("Length = %d\n", len(mySlice))
	fmt.Printf("capacity = %d\n", cap(mySlice))

	mySlice2 := make([]int, 5)

	fmt.Printf("mySlice2 = %v\n", mySlice2)
	fmt.Printf("length = %d\n", len(mySlice2))
	fmt.Printf("capacity = %d\n", cap(mySlice2))
}