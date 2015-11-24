# chancast

Broadcasting using channels, rather than `sync.Cond`, because `sync.Cond.Wait` sucks.

### Quick Example

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
