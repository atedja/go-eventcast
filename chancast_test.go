package chancast

import (
	"sync"
	"testing"
	"time"
)

func TestEmptyBroadcast(t *testing.T) {
	Init("hello")
	for i := 0; i < 10; i++ {
		Broadcast("hello")
	}
}

func TestSingleBroadcast(t *testing.T) {
	Init("hello")
	go func() {
		time.Sleep(100 * time.Millisecond)
		Broadcast("hello")
	}()

	<-Wait("hello")
}

func TestMultipleBroadcast(t *testing.T) {
	Init("hello")

	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			<-Wait("hello")
			wg.Done()
		}()
	}

	time.Sleep(100 * time.Millisecond)
	Broadcast("hello")

	wg.Wait()
}
