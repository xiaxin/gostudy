package test

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncMap(test *testing.T) {
	var hash sync.Map

	hash.Store("key", "val")
	hash.Store("k1", "v1")

	{
		if val, ok := hash.Load("key"); ok {
			fmt.Println(val)
		}
	}

	{
		val, loaded := hash.LoadOrStore("key", "value")

		fmt.Println(val, loaded)
	}

	{
		val, loaded := hash.LoadOrStore("key", "value")

		fmt.Println(val, loaded)
	}

	{
		val, loaded := hash.LoadAndDelete("key")

		fmt.Println(val, loaded)
	}

	{
		hash.Range(func(key, val interface{}) bool {
			fmt.Println(key, val)
			return true
		})
	}
}
