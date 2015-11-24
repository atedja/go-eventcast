# chancast

Broadcasting using channels, rather than `sync.Cond`, because `sync.Cond.Wait` sucks.

### Quick Example

```go
  chancast.Init("we are closed")

  // spawn workers
  go func() {
    for {
      select {
      case <-chancast.Wait("we are closed"):
        return
      default:
        // do other things
      }
    }
  }()

  // sometime later..
  chancast.Broadcast("we are closed")
```
