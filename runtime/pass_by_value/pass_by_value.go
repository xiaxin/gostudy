package main

import "fmt"

// 函数 pass-by-value

func main() {

	{
		var i int = 10
		fmt.Printf("src:%x\n", &i)
		FnInt(i)
	}

	{ // slice
		var ints []int = []int{1, 2, 3, 4, 5}
		fmt.Printf("src:%x\n", &ints)
		fmt.Printf("srcVal:%v\n", ints)
		FnSlice(ints)
		fmt.Printf("srcVal:%v\n", ints)
	}
}

func FnInt(i int) {
	fmt.Printf("dst:%x\n", &i)
}

func FnSlice(ints []int) {
	ints[2] = 100

	fmt.Printf("dst:%x\n", &ints)
	fmt.Printf("dstVal:%v\n", ints)
}
