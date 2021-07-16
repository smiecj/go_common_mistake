package go_common_mistake

import (
	"log"
	"testing"
)

func TestInitSlice(t *testing.T) {
	// emptySlice := make([]int, 0)
	emptySlice := make([]int, 0, 0)

	sliceWithIlegalCap := make([]int, 0, 1)

	log.Printf("%v %v", emptySlice, sliceWithIlegalCap)
}

// todo: add init len or cap compare
