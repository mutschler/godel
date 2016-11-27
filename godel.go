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
  if deviceuid != "" {
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

  // fmt.Println(message)

  mac := hmac.New(sha1.New, []byte("aPLFAjyUusVPHgcgvlAxihthmRaiuqCjBsRCPLan"))
	mac.Write([]byte(message))
	expectedMAC := mac.Sum(nil)
  hash := fmt.Sprintf("%x", expectedMAC)
  return hash, timestamp
}

//sendRequest handles all sorts of request to the jodel api
func (g *Godel) sendRequest(endpoint string, parameters interface{}, payload interface{}, method string) *grequests.Response {
  url := fmt.Sprintf("https://api.go-tellm.com/api/v2/%s", endpoint)

  if(parameters != nil) {
    params, _ := query.Values(parameters)
    url = url + "?" + strings.ToLower(params.Encode())
  }

  payloadString, err := json.Marshal(payload)
  if err != nil {
      fmt.Println("error ", err)
  }

  //sign the request using stringified request and hmac
  hmac, timestamp := signRequest(string(payloadString), endpoint, g.AccessToken, method)

  ro := &grequests.RequestOptions{
    UserAgent: "Jodel/4.28.1 Dalvik/2.1.0 (Linux; U; Android 5.1.1; D6503 Build/23.4.A.1.232)",
    DisableCompression: false,
    Headers: map[string]string{"Connection": "keep-alive","Content-Type": "application/json; charset=UTF-8","X-Timestamp": timestamp, "X-Client-Type": "android_4.28.1", "X-Api-Version": "0.2", "X-Authorization": "HMAC " + hmac,"Authorization": "Bearer " + g.AccessToken},
    JSON: payload,
  }

  var resp *grequests.Response

  switch method {
    case "POST":
      resp, err = grequests.Post(url, ro)
    case "GET":
      resp, err = grequests.Get(url, ro)
    case "PUT":
      resp, err = grequests.Put(url, ro)
    case "DELETE":
      resp, err = grequests.Put(url, ro)
    default:
      resp, err = grequests.Get(url, ro)
  }

  if err != nil {
      fmt.Println("Unable to make request: ", err)
  }

  fmt.Println(resp.StatusCode)
  fmt.Println(resp.String())
  fmt.Println(resp.Header)
  return resp
}

func (g *Godel) get(endpoint string, parameters interface{}) *grequests.Response {
  return g.sendRequest(endpoint, parameters, nil, "GET")
}

func (g *Godel) put(endpoint string, payload interface{}) *grequests.Response {
  return g.sendRequest(endpoint, nil, payload, "PUT")
}

func (g *Godel) delete(endpoint string, payload interface{}) *grequests.Response {
  return g.sendRequest(endpoint, nil, payload, "DELETE")
}

// Internal post function
func (g *Godel) post(endpoint string, payload interface{}) *grequests.Response {
  return g.sendRequest(endpoint, nil, payload, "POST")
}


//GetRequestToken sends user location and deviceUID to get a request token
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
  x := LoginResponse{}
  response.JSON(&x)

  if(x.AccessToken != "" && response.Ok) {
    g.Location = location
    g.AccessToken = x.AccessToken
    g.DistinctID = x.DistinctID
    g.RefreshToken = x.RefreshToken
    return nil
  }
  return nil
}

//GetNewAccessToken refresh the access token
func (g *Godel) GetNewAccessToken() {
  data := NewAccessTokenRequest{
    clientID,
    g.DistinctID,
    g.RefreshToken,
  }
  response := g.post("users/refreshToken", data)

  token := NewAccessTokenResonse{}
  response.JSON(&token)

  if response.Ok {
    g.AccessToken = token.AccessToken
  }
}

//GetMostRecentPosts get the most recent posts for the current location
func (g *Godel) GetMostRecentPosts() []SinglePostResponse {
  response := g.get("posts/location/", nil)
  posts := LatestPostsResponse{}
  response.JSON(&posts)

  return posts.Posts
}

//GetMostPopularPosts get most popular posts for the current location
func (g *Godel) GetMostPopularPosts() []SinglePostResponse  {
  response := g.get("posts/location/popular", nil)
  posts := LatestPostsResponse{}
  response.JSON(&posts)

  return posts.Posts
}

//GetMostDiscussedPosts get most discussed posts for the current location
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

//GetMyPinnedPosts returns a list of your pinned posts
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
  response := g.get("posts/mine/replies", nil)
  post := LatestPostsResponse{}
  response.JSON(&post)

  return post.Posts
}

//GetMyVotedPosts returns a list of posts you've replied to
func (g *Godel) GetMyVotedPosts() []SinglePostResponse  {
  response := g.get("posts/mine/voted", nil)
  post := LatestPostsResponse{}
  response.JSON(&post)

  return post.Posts
}

//GetMyPostsCombo returns a list of your last posts, your votes and comments
func (g *Godel) GetMyPostsCombo() ComboPosts  {
  response := g.get("posts/mine/combo", nil)
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
func (g *Godel) PinPost(postid string) bool  {
  response := g.put("posts/"+postid+"/pin", nil)
return response.Ok
}

//UnpinPost upvote a selected Post
func (g *Godel) UnpinPost(postid string) bool  {
  response := g.put("posts/"+postid+"/unpin", nil)
  return response.Ok
}

//SendPost creates a new Jodel Post
func (g *Godel) SendPost(message, color string) bool {
  payload := NewPost{
    color,
    g.Location,
    message,
  }

  response := g.post("posts/", payload)
  return response.Ok
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
