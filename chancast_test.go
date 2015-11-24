package chancast

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
	"time"
)

func TestNoListener(t *testing.T) {
	for i := 0; i < 10; i++ {
		Broadcast("hello")
	}
}

func TestSingleListener(t *testing.T) {
	go func() {
		time.Sleep(100 * time.Millisecond)
		Broadcast("hello")
	}()

	<-Listen("hello")
}

func TestMultipleListeners(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			<-Listen("hello")
			wg.Done()
		}()
	}

	time.Sleep(100 * time.Millisecond)
	Broadcast("hello")

	wg.Wait()
}

func TestBroadcastWithValue(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			data := <-Listen("hello")
			assert.Equal(t, "value", data.(string))
			wg.Done()
		}()
	}

	time.Sleep(100 * time.Millisecond)
	BroadcastWithValue("hello", "value")

	wg.Wait()
}

func TestHeyHoo(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(6)
	for i := 0; i < 2; i++ {
		go func() {
			hey := Listen("hey")
			hoo := Listen("hoo")
			done := Listen("done")
			for {
				select {
				case _, ok := <-hey:
					if ok {
						wg.Done()
					}
				case _, ok := <-hoo:
					if ok {
						wg.Done()
					}
				case <-done:
					wg.Done()
					return
				}
			}
		}()
	}

	time.Sleep(100 * time.Millisecond)
	Broadcast("hey")
	Broadcast("hoo")
	time.Sleep(100 * time.Millisecond)
	Broadcast("done")
	wg.Wait()
}
