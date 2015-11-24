# chancast

Broadcasting using channels, rather than `sync.Cond`.

### Quick Example

```go
  chancast.Init("we are closed")

  // spawn workers
  go func() {
    select {
    case <-chancast.Wait("we are closed"):
      return
    default:
      // do other things
    }
  }()

  // sometime later..
  chancast.Broadcast("we are closed")
```
