package main

import "fmt"

func main() {

	list := new([]int)

	*list = append(*list, 1)

	fmt.Println(list)

	// a := []int{1}

	// b := []int{11, 22}

	// fmt.Printf("a[%p] %v %d %d\n", &a, a, len(a), cap(a))
	// fmt.Printf("b[%p] %v %d %d\n", &b, b, len(b), cap(b))

	// s := append(a, b...)

	// fmt.Printf("s[%p] %v %d %d\n", &s, s, len(s), cap(s))

	// s[2] = 99
	// fmt.Printf("new a[%p] %v %d %d\n", &a, a, len(a), cap(a))
	// fmt.Printf("new s[%p] %v %d %d\n", &s, s, len(s), cap(s))

	// {
	// 	x := []int{0, 0, 0}

	// 	for i := 0; i < 5; i++ {
	// 		x = append(x, i)
	// 		fmt.Println(x, len(x), cap(x))
	// 	}

	// }
}
