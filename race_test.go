package go_common_mistake

import (
	"fmt"
	"testing"
)

func TestMapRace(t *testing.T) {
	done := make(chan bool)
	m := make(map[string]string)
	m["name"] = "world"
	go func() {
		m["name"] = "data race"
		done <- true
	}()
	fmt.Println("Hello,", m["name"])
	<-done
}
