package go_common_mistake

import (
	"fmt"
	"sync"
	"testing"
)

func TestRoutineRace(t *testing.T) {
	ints := []int{1, 2, 3}
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(ints))

	for _, i := range ints {
		go func() {
			fmt.Printf("%v\n", i)
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
}

func TestRoutine(t *testing.T) {
	ints := []int{1, 2, 3}
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(ints))

	for _, i := range ints {
		go func(i int) {
			fmt.Printf("%v\n", i)
			waitGroup.Done()
		}(i)
	}
	waitGroup.Wait()
}
