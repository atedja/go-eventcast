# eventcast

Simple event broadcasting using channels. No locking needed.

### Examples

#### Signaling arbitrary number of workers

```go
  // spawn workers
  for i := 0; i < 100; i++ {
    go func() {
      closed := eventcast.Listen("we are closed")
      for {
        select {
        case <-closed:
          return
        default:
          // do other things
        }
      }
    }()
  }

  // somewhere, sometime later..
  eventcast.Broadcast("we are closed")
```

#### Broadcasting a value to multiple listeners

```go
  // goroutines waiting for some result or timeout.
  for i := 0; i < 10; i++ {
    go func() {
      select {
      case value := <-eventcast.Listen("result"):
        // do something with value
      case <-time.After(1 * time.Second):
        // timeout!
        break
      }
    }()
  }

  // some worker
  go func() {
    // doing something...

    value := "result of processing some data"
    eventcast.BroadcastWithValue("result", value)

    // continue doing more
  }()
```
