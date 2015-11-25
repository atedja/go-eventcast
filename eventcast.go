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

// Broadcast, unblock all channels.
//
func Broadcast(key string) {
	BroadcastWithValue(key, nil)
}

// Broadcast with a value.
//
func BroadcastWithValue(key string, value interface{}) {
	channels.Lock()
	defer channels.Unlock()
	chans := channels.list[key]
	if chans != nil {
		for _, ch := range chans {
			ch <- value
			close(ch)
		}
		channels.list[key] = nil
	}
}

// Listener. Retrieve a channel that listens to a broadcast event.
// This method creates a one-time use channel, everytime, and a Broadcast() will close and remove them.
//
func Listen(key string) chan interface{} {
	channels.Lock()
	defer channels.Unlock()
	out := make(chan interface{}, 1)
	if channels.list[key] == nil {
		channels.list[key] = make([]chan interface{}, 0, 8)
	}
	channels.list[key] = append(channels.list[key], out)
	return out
}
