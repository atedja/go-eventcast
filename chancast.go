package chancast

import (
	"sync"
)

type channelList struct {
	sync.Mutex
	list map[string][]chan bool
}

var channels = channelList{
	list: make(map[string][]chan bool),
}

// Broadcasts, unblock all channels.
//
func Broadcast(key string) {
	channels.Lock()
	defer channels.Unlock()
	chans := channels.list[key]
	if chans != nil {
		count := len(chans)
		for i := 0; i < count; i++ {
			chans[i] <- true
			close(chans[i])
		}
		channels.list[key] = nil
	}
}

// Listener. Retrieve a channel that listens to a broadcast.
// This method creates a one-time use channel, everytime, and a Broadcast will remove it.
//
func Listen(key string) chan bool {
	channels.Lock()
	defer channels.Unlock()
	out := make(chan bool, 1)
	if channels.list[key] == nil {
		channels.list[key] = make([]chan bool, 0)
	}
	channels.list[key] = append(channels.list[key], out)
	return out
}
