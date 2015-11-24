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
	for i := len(chans) - 1; i >= 0; i-- {
		chans[i] <- true
		chans = append(chans[:i], chans[i+1:]...)
	}
	channels.list[key] = chans
}

// Listening...
//
func Wait(key string) <-chan bool {
	channels.Lock()
	defer channels.Unlock()
	out := make(chan bool)
	channels.list[key] = append(channels.list[key], out)
	return out
}

// Must be called before using the key.
//
func Init(key string) {
	channels.Lock()
	defer channels.Unlock()
	if channels.list[key] == nil {
		channels.list[key] = make([]chan bool, 0)
	}
}
