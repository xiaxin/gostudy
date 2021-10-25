package array

import (
	"log"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	log.Println(5, binarySearch(nums, 5))
	log.Println(0, binarySearch(nums, 0))
	log.Println(9, binarySearch(nums, 9))
	log.Println(10, binarySearch(nums, 10))
}

func TestBinarySearchBoard(t *testing.T) {
	{
		nums := []int{0, 1, 1, 1, 2}

		log.Println(1, binarySearchLeftBoard(nums, 1))
		log.Println(10, binarySearchLeftBoard(nums, 10))
	}

	{
		nums := []int{0, 1, 1, 1, 2}

		log.Println(1, binarySearchRightBoard(nums, 1))
		log.Println(10, binarySearchRightBoard(nums, 10))
	}

}
