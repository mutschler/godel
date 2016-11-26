package godel

import (
  "encoding/json"
  "github.com/levigross/grequests"
  "fmt"
  "math/rand"
  "time"
  "crypto/hmac"
  "crypto/sha1"
  "github.com/google/go-querystring/query"
  "strings"
)

const clientID = "81e8a76e-1e02-4d17-9ba0-8a7020261b26";
const apiBaseURL = "htts://go-tellm.de/api/v2/"

//RED color
const RED = "DD5F5F"
//BLUE color
const BLUE = "DD5F5F"
//ORANGE color
const ORANGE = "FF9908"
//GREEN color
const GREEN = "9EC41C"
//LIMEGREEN color
const LIMEGREEN = "8ABDB0"
//YELLOW color
const YELLOW = "FFBA00"
//Pick a random color
const RANDOM = "FFFFFF"

const runes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

//Godel struct
type Godel struct {
  Key string
  DeviceUID string
  AccessToken string
  RefreshToken string
  DistinctID string
  Location LocationResponse
}

//NewClient create a new Godel Client, if no deviceuid is given a random one will be created
func NewClient(deviceuid string) (_ *Godel) {
  uid := NewDeviceUID()
  //TODO: fix this: if empty generate new Device!!!
  if (deviceuid != "") {
    uid = deviceuid
  }
  return &Godel{
    Key: "aPLFAjyUusVPHgcgvlAxihthmRaiuqCjBsRCPLan",
    DeviceUID: uid,
  }
}

//NewDeviceUID creates a new DeviceUID use this, or create your own 63 char rand string
func NewDeviceUID() (uid string) {
  b := make([]byte, 63) // create ne 63 letter string
  for i := range b {
    b[i] = runes[rand.Int63() % int64(len(runes))]
  }
  return string(b)
}

//this handles request signing...
func signRequest(body, url, auth, method string) (hmacresult, ts string) {
  timestamp := time.Now().UTC().Format(time.RFC3339)
  parameters := ""
  message := method + "%" + "api.go-tellm.com" + "%" + "443" + "%/api/v2/" + url + "%" + auth + "%" + timestamp + "%" + parameters + "%" + body

  fmt.Println(message)

  mac := hmac.New(sha1.New, []byte("aPLFAjyUusVPHgcgvlAxihthmRaiuqCjBsRCPLan"))
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)
  hash := fmt.Sprintf("%x", expectedMAC)
  return hash, timestamp
  // return base64.StdEncoding.EncodeToString(expectedMAC), timestamp
  // return "1f8b08200519d33ee62f9c7e50893e43dac8cdf7", timestamp
}

// TODO: requests could be merged to one func which just calls get, put, post depending on modes less redundant! and all requests get signed
// Internal get function
// func (g *Godel) get(endpoint string) *grequests.Response {
//   ro := &grequests.RequestOptions{
//     UserAgent: "Jodel/4.28.1 Dalvik/2.1.0 (Linux; U; Android 6.0.1; Nexus 5 Build/MMB29V)",
//     Headers: map[string]string{"Connection": "keep-alive","Content-Type": "application/json; charset=UTF-8","Authorization": "Bearer " + g.AccessToken},
//     DisableCompression: false,
//   }
//   url := fmt.Sprintf("https://api.go-tellm.com/api/v2/%s", endpoint)
//   resp, err := grequests.Get(url, ro)
//   // You can modify the request by passing an optional RequestOptions struct
//
//   if err != nil {
//       fmt.Println("Unable to make request: ", err)
//   }
//   // fmt.Println(resp.String())
//   fmt.Println(resp.StatusCode)
//   fmt.Println(resp.String())
//   fmt.Println(resp.Header)
//   return resp
// }
func (g *Godel) get(endpoint string, payload interface{}) *grequests.Response {
  ro := &grequests.RequestOptions{
    UserAgent: "Jodel/4.28.1 Dalvik/2.1.0 (Linux; U; Android 6.0.1; Nexus 5 Build/MMB29V)",
    Headers: map[string]string{"Connection": "keep-alive","Content-Type": "application/json; charset=UTF-8","Authorization": "Bearer " + g.AccessToken},
    DisableCompression: false,
    JSON: payload,
  }
  url := fmt.Sprintf("https://api.go-tellm.com/api/v2/%s", endpoint)

  // if (queryParams) {
	// 		queryParams = queryParams ? queryParams : {};
	// 		queryParams = humps.decamelizeKeys(queryParams);
  //
	// 		qs = querystring.stringify(queryParams);
	// 		if (qs && qs.length > 0) {
	// 			if (!url.endsWith("?"))
	// 				url += "?";
	// 			url += qs;
	// 		}
	// 	}

  if(payload != nil) {
    params, _ := query.Values(payload)
    url = url + "?" + strings.ToLower(params.Encode())
  }

  fmt.Println(url)

  resp, err := grequests.Get(url, ro)
  // You can modify the request by passing an optional RequestOptions struct

  if err != nil {
      fmt.Println("Unable to make request: ", err)
  }
  // fmt.Println(resp.String())
  fmt.Println(resp.StatusCode)
  fmt.Println(resp.String())
  fmt.Println(resp.Header)
  return resp
}

func (g *Godel) put(endpoint string, payload interface{}) *grequests.Response {
  ro := &grequests.RequestOptions{
    UserAgent: "Jodel/4.28.1 Dalvik/2.1.0 (Linux; U; Android 6.0.1; Nexus 5 Build/MMB29V)",
    Headers: map[string]string{"Connection": "keep-alive","Content-Type": "application/json; charset=UTF-8","Authorization": "Bearer " + g.AccessToken},
    DisableCompression: false,
    JSON: payload,
  }
  url := fmt.Sprintf("https://api.go-tellm.com/api/v2/%s", endpoint)
  fmt.Println(url)
  resp, err := grequests.Put(url, ro)
  // You can modify the request by passing an optional RequestOptions struct

  if err != nil {
      fmt.Println("Unable to make request: ", err)
  }

  fmt.Println(resp.StatusCode)
  fmt.Println(resp.String())
  fmt.Println(resp.Header)

  fmt.Println(resp.String())
  return resp
}

func (g *Godel) delete(endpoint string, payload interface{}) *grequests.Response {
  ro := &grequests.RequestOptions{
    UserAgent: "Jodel/4.28.1 Dalvik/2.1.0 (Linux; U; Android 6.0.1; Nexus 5 Build/MMB29V)",
    Headers: map[string]string{"Connection": "keep-alive","Content-Type": "application/json; charset=UTF-8","Authorization": "Bearer " + g.AccessToken},
    DisableCompression: false,
    JSON: payload,
  }
  url := fmt.Sprintf("https://api.go-tellm.com/api/v2/%s", endpoint)
  fmt.Println(url)
  resp, err := grequests.Delete(url, ro)
  // You can modify the request by passing an optional RequestOptions struct

  if err != nil {
      fmt.Println("Unable to make request: ", err)
  }

  fmt.Println(resp.StatusCode)
  fmt.Println(resp.String())
  fmt.Println(resp.Header)

  fmt.Println(resp.String())
  return resp
}

// Internal post function
func (g *Godel) post(endpoint string, payload interface{}) *grequests.Response {

  out, err := json.Marshal(payload)
  if err != nil {
      panic (err)
  }

  fmt.Println(g.AccessToken)

  hmac, timestamp := signRequest(string(out), endpoint, g.AccessToken, "POST")

  ro := &grequests.RequestOptions{
    UserAgent: "Jodel/4.28.1 Dalvik/2.1.0 (Linux; U; Android 5.1.1; D6503 Build/23.4.A.1.232)",
    DisableCompression: false,
    Headers: map[string]string{"Connection": "keep-alive","Content-Type": "application/json; charset=UTF-8","X-Timestamp": timestamp, "X-Client-Type": "android_4.28.1", "X-Api-Version": "0.2", "X-Authorization": "HMAC " + hmac,"Authorization": "Bearer " + g.AccessToken},
    JSON: payload,
  }

  // resp, err := grequests.Post("http://httpbin.org/post", ro)
  url := fmt.Sprintf("https://api.go-tellm.com/api/v2/%s", endpoint)
  resp, err := grequests.Post(url, ro)
  if err != nil {
      fmt.Println("Unable to make request: ", err)
  }

  fmt.Println(resp.StatusCode)
  fmt.Println(resp.String())
  fmt.Println(resp.Header)
  //
  // x := LoginResponse{}
  // resp.JSON(&x)
  //
  // resp.DownloadToFile("test.txt")

  return resp
}


//GetRequestToken is the login function
//TODO: rename this
func (g *Godel) GetRequestToken(city string, country string, lat float64, lng float64) (err error) {

  location := LocationResponse{
    19,
    city,
    country,
    city,
    LocationCoordinates{
      lat,
      lng,
    },
  }

  login := LoginPayload{
    clientID,
    g.DeviceUID,
    location,
  }

  response := g.post("users/", login)
  // result := g.post("/users/", payload)
  x := LoginResponse{}
  response.JSON(&x)
  fmt.Println(x)
  g.Location = location
  g.AccessToken = x.AccessToken
  g.DistinctID = x.DistinctID
  g.RefreshToken = x.RefreshToken
  if(x.AccessToken != "" && response.Ok) {
    return nil
  }
  return nil
}

//GetNewAccessToken refreshs the access token
func (g *Godel) GetNewAccessToken() {
  data := NewAccessTokenRequest{
    clientID,
    g.DistinctID,
    g.RefreshToken,
  }
  response := g.post("users/refreshToken", data)

  token := NewAccessTokenResonse{}
  response.JSON(&token)

  g.AccessToken = token.AccessToken
}

//NewDeviceUID returns a new DeviceUID. Use this or create your own...
func (g *Godel) NewDeviceUID() string {
  return NewDeviceUID()
}

//GetMostRecentPosts get new Posts for current location
func (g *Godel) GetMostRecentPosts() []SinglePostResponse {
  response := g.get("posts/location/", nil)
  posts := LatestPostsResponse{}
  response.JSON(&posts)

  return posts.Posts
}

//GetMostPopularPosts get top posts for current location
func (g *Godel) GetMostPopularPosts() []SinglePostResponse  {
  response := g.get("posts/location/popular", nil)
  posts := LatestPostsResponse{}
  response.JSON(&posts)

  return posts.Posts
}

//GetMostDiscussedPosts get top posts for current location
func (g *Godel) GetMostDiscussedPosts(pagination interface{}) []SinglePostResponse  {
  response := g.get("posts/location/discussed", pagination)
  posts := LatestPostsResponse{}
  response.JSON(&posts)

  return posts.Posts
}

//GetMyPosts returns a list of your own posts
func (g *Godel) GetMyPosts() []SinglePostResponse {
  response := g.get("posts/mine", nil)
  posts := LatestPostsResponse{}
  response.JSON(&posts)

  return posts.Posts
}

//GetMyPinnedPosts returns a list of your own posts
func (g *Godel) GetMyPinnedPosts() []SinglePostResponse {
  response := g.get("posts/mine/pinned", nil)
  posts := LatestPostsResponse{}
  response.JSON(&posts)

  return posts.Posts
}

//GetMyPopularPosts returns a list of your own posts
func (g *Godel) GetMyPopularPosts() []SinglePostResponse {
  response := g.get("posts/mine/popular", nil)
  posts := LatestPostsResponse{}
  response.JSON(&posts)

  return posts.Posts
}

//GetMyMostDiscussedPosts returns a list of your own posts
func (g *Godel) GetMyMostDiscussedPosts() []SinglePostResponse {
  response := g.get("posts/mine/discussed", nil)
  posts := LatestPostsResponse{}
  response.JSON(&posts)

  return posts.Posts
}


//GetMyRepliedPosts returns a list of posts you've replied to
func (g *Godel) GetMyRepliedPosts() []SinglePostResponse  {
  response := g.get("posts/mine/replies/", nil)
  post := LatestPostsResponse{}
  response.JSON(&post)

  return post.Posts
}

//GetMyVotedPosts returns a list of posts you've replied to
func (g *Godel) GetMyVotedPosts() []SinglePostResponse  {
  response := g.get("posts/mine/voted/", nil)
  post := LatestPostsResponse{}
  response.JSON(&post)

  return post.Posts
}

//GetMyPostsCombo returns a list of your last posts, your votes and comments
func (g *Godel) GetMyPostsCombo() ComboPosts  {
  response := g.get("posts/mine/combo/", nil)
  post := ComboPosts{}
  response.JSON(&post)

  return post
}


//GetKarma returns a list of your own posts
func (g *Godel) GetKarma() KarmaResponse {
  response := g.get("users/karma", nil)
  karma := KarmaResponse{}
  response.JSON(&karma)

  return karma
}

//GetPost returns a single post
func (g *Godel) GetPost(postid string) SinglePostResponse {
  response := g.get("posts/" + postid, nil)
  posts := SinglePostResponse{}
  response.JSON(&posts)

  return posts
}

//DeletePost is used to delete a post
func (g *Godel) DeletePost(postid string) bool {
  response := g.delete("posts/"+ postid, nil)
  return response.Ok
}

//UpvotePost upvote a selected Post
func (g *Godel) UpvotePost(postid string) SinglePostResponse  {
  response := g.put("posts/"+postid+"/upvote", nil)
  post := VoteResponse{}
  response.JSON(&post)

  return post.Post
}

//DownvotePost upvote a selected Post
func (g *Godel) DownvotePost(postid string) SinglePostResponse  {
  response := g.put("posts/"+postid+"/downvote", nil)
  post := VoteResponse{}
  response.JSON(&post)

  return post.Post
}

//PinPost upvote a selected Post
func (g *Godel) PinPost(postid string) SinglePostResponse  {
  response := g.put("posts/"+postid+"/pin", nil)
  post := VoteResponse{}
  response.JSON(&post)

  return post.Post
}

//UnpinPost upvote a selected Post
func (g *Godel) UnpinPost(postid string) SinglePostResponse  {
  response := g.put("posts/"+postid+"/unpin", nil)
  post := VoteResponse{}
  response.JSON(&post)

  return post.Post
}

//SendPost creates a new Jodel Post
func (g *Godel) SendPost(message, color string) {
  payload := NewPost{
    color,
    g.Location,
    message,
  }

  response := g.post("posts/", payload)
  fmt.Println(response.String())
}

//SendReply allows you to reply to a post
func (g *Godel) SendReply(postid, message, color string) SinglePostResponse {
  reply := ReplyToPost{
    postid,
    color,
    g.Location,
    message,
  }
  response := g.post("posts", reply)
  posts := SinglePostResponse{}
  response.JSON(&posts)

  return posts
}

//SendUserLocation can be used to change the location withouth the need to relogin
func (g *Godel) SendUserLocation(city string, country string, lat, lng float64) LocationResponse {

  loc := LocationResponse{
    19,
    city,
    country,
    city,
    LocationCoordinates{
      lat,
      lng,
    },
  }

  newLoc := &SendLocationRequest{
    loc,
  }
  response := g.put("users/location", newLoc)
  g.Location = loc
  fmt.Println(response.String())
  return loc
}
