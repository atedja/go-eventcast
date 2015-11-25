package eventcast

import (
	"sync"
)

var channels = struct {
	sync.Mutex
	list map[string][]chan interface{}
}{
	list: make(map[string][]chan interface{}),
}

// Broadcast an event.
//
func Broadcast(event string) {
	BroadcastWithValue(event, nil)
}

// Broadcast an event with a value.
//
func BroadcastWithValue(event string, value interface{}) {
	channels.Lock()
	defer channels.Unlock()
	chans := channels.list[event]
	if chans != nil {
		for _, ch := range chans {
			ch <- value
			close(ch)
		}
		channels.list[event] = nil
	}
}

// Listener. Retrieve a channel that listens to a broadcast event.
// This method creates a one-time use channel, everytime, and a Broadcast()
// will push data and close them.
//
func Listen(event string) chan interface{} {
	channels.Lock()
	defer channels.Unlock()
	out := make(chan interface{}, 1)
	if channels.list[event] == nil {
		channels.list[event] = make([]chan interface{}, 0, 8)
	}
	channels.list[event] = append(channels.list[event], out)
	return out
}
