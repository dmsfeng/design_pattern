package singleton

import (
	"sync"

	"testing"
)

const parCount = 100

func TestSingleton(t *testing.T) {
	ins1 := getInstance()
	ins2 := getInstance()
	if ins1 != ins2 {
		t.Fatal("instance is not equal")
	}
}

func TestParallelSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]*singleton{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			instances[index] = getInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance is not equal")
		}
	}
}
