package go_common_mistake

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
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

func TestSyncMapRace(t *testing.T) {
	done := make(chan bool)
	m := sync.Map{}
	m.Store("name", "world")
	go func() {
		m.Store("name", "data race")
		done <- true
	}()
	v, _ := m.Load("name")
	fmt.Println("Hello,", v)
	<-done
}

func randomDuration() time.Duration {
	rand.Seed(time.Now().Unix())
	return time.Duration(rand.Int63n(1e9))
}

func TestAfterFuncRace(t *testing.T) {
	start := time.Now()
	var timer *time.Timer
	timer = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		timer.Reset(randomDuration())
	})
	time.Sleep(5 * time.Second)
}

func TestAfterFunc(t *testing.T) {
	start := time.Now()
	var timer *time.Timer
	needResetTimerChan := make(chan bool)
	timer = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		needResetTimerChan <- true
	})
	for time.Since(start) < 5*time.Second {
		<-needResetTimerChan
		timer.Reset(randomDuration())
	}
}
