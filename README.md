# godel: a simple jodel API in go


```go get github.com/mutschler/godel```


Godel supports all basic Jodel features like, getting posts, up/downvoting, post, reply etc. For Detailed usage please see https://godoc.org/github.com/mutschler/godel

godel hast the ability to print some output to stdout (usefull mostly for debugging) to activate just set `client.Debug = true` also there is the ability to log/dump every response to a json file inside a subdirectory to use this, you'll have to provide a base direcotry like this `client.DownloadDir = "mydir"`.

#### example usage

```
  import "github.com/mutschler/godel"
  import "fmt"

  client = godel.NewClient(godel.NewDeviceUID())

  //turn on debug output
  client.Debug = true

  //provide a base dir for downloading data
  client.DownloadDir = "/path/to/my/dl_dir/"

  client.GetRequestToken("Berlin", "DE", 52.520007, 13.404954)  

  //fetch most popular posts and print message  
  for _, post := range g.GetMostPopularPosts() {  
    fmt.Println(post.Message)  
  }
```
