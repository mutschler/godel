# WIP

## godel: a simple jodel API in go


```go get github.com/mutschler/godel```

### Currently unstable and very likely to change a lot so use at your own risk

#### example

```
  import "github.com/mutschler/godel"
  import "fmt"

  client = godel.NewClient(godel.NewDeviceUID())  
  client.GetRequestToken("Berlin", "DE", 52.520007, 13.404954)  

  //fetch most popular posts and print message  
  for _, post := range g.GetMostPopularPosts() {  
    fmt.Println(post.Message)  
  }
```
