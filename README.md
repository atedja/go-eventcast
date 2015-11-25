# chancast

Broadcasting using channels, rather than `sync.Cond`, because `sync.Cond.Wait` sucks.

### Examples


#### Signaling arbitrary number of workers

```go
  // spawn workers
  for i := 0; i < 100; i++ {
    go func() {
      closed := chancast.Listen("we are closed")
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
  chancast.Broadcast("we are closed")
```

#### Broadcasting a value to multiple listeners

```go
  // goroutines waiting for some result
  for i := 0; i < 10; i++ {
    go func() {
      select {
      case value := <-chancast.Listen("result"):
        // do something with result
      }
    }()
  }

  // some worker
  go func() {
    // doing something
    value := "result of processing some data"

    chancast.BroadcastWithValue("result", value)

    // continue doing more
  }()
```
