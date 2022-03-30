package test

import (
	"sync"
	"testing"
)

func TestOnce(t *testing.T) {
	var once sync.Once
	onceBody := func() {
		t.Log("Only once")
	}
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	for i := 0; i < 10; i++ {
		<-done
	}
}
