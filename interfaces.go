package godel

// LoginResponse holds the login response from the Server
type LoginResponse struct {
  AccessToken string `json:"access_token"`
  RefreshToken string `json:"refresh_token"`
  TokenType string `json:"token_type"`
  ExpiresIn int64 `json:"expires_in"`
  ExpirationDate int64 `json:"expiration_date"`
  DistinctID string `json:"distinct_id"`
  Returning bool `json:"returning"`
}

//LatestPostsResponse ...
type LatestPostsResponse struct {
  Max int64 `json:"max"`
  Posts []SinglePostResponse `json:"posts"`
}

//SingleComment holds one comment of a post
type SingleComment struct {
  Color string `json:"color"` // HEX Color of the post
  CreatedAt string `json:"created_at"` // ISO Timestamp
  Discovered int64 `json:"discovered"` // unknown
  DiscoveredBy int64 `json:"discovered_by"` // unknown
  Distance int64 `json:"distance"` // distance to current location
  Location Location `json:"location"` //Location struct
  Message string `json:"message"` // the message
  ParentCreator int64 `json:"parent_creator"` // 1 if reply is from OJ else 0
  PostID string `json:"post_id"` // the unique postID
  PostOwn string `json:"post_own"` // own if own post else friend
  UpdatedAt string `json:"updated_at"` // ISO Timestamp
  UserHandle string `json:"user_handle"` // the user handle of the poster
  VoteCount int64 `json:"vote_count"` // how many upvotes
}

// SinglePostResponse holds the response for one single Post
type SinglePostResponse struct {
  ChildCount int64 `json:"child_count"` // count of replys
  Children []SingleComment `json:"children"` // list of replys
  Color string `json:"color"` // HEX Color of the post
  CreatedAt string `json:"created_at"` // ISO Timestamp
  Discovered int64 `json:"discovered"` // unknown
  DiscoveredBy int64 `json:"discovered_by"` // unknown
  Distance int64 `json:"distance"` // distance to current location
  GotThanks bool `json:"got_thanks"` // bool
  Location Location `json:"location"` // Location
  Message string `json:"message"` // the message
  NotificationsEnabled bool `json:"notifications_enabled"` // notifications for new replys
  PinCount int64 `json:"pin_count"` // how many pins
  PostID string `json:"post_id"` // the unique postID
  PostOwn string `json:"post_own"` // own if own post else friend
  UpdatedAt string `json:"updated_at"` // ISO Timestamp
  UserHandle string `json:"user_handle"` // the user handle of the poster
  VoteCount int64 `json:"vote_count"` // how many upvotes
  Voted string `json:"voted"` // if voted then up or down
}

// SendLocationRequest is used to set a new Location for a user
type SendLocationRequest struct {
  Location Location `json:"location"` // Location struct
}

// NewPost can be used to create a new Post
type NewPost struct {
  Color string `json:"color"` // HEX Color of the post
  Location Location `json:"location"` // Location struct
  Message string `json:"message"` // the message
}

//Coordinates hold Latitude and Longitude of the current location
type Coordinates struct {
  Lat float64 `json:"lat"` // Latitude
  Lng float64 `json:"lng"` // Longitude
}

//Location Holds the users current Location
type Location struct {
  LocAccuracy float64 `json:"loc_accuracy"` // accuracy this is randomly generated
  City string `json:"city"` // Name of the City
  Country string `json:"country"` // ISO Country Code (DE for Germany)
  Name string `json:"name"` // Name of the City
  LocCoordinates Coordinates `json:"loc_coordinates"` // Coordinates struct
}

//LoginPayload is used to log the user in
type LoginPayload struct {
  ClientID string `json:"client_id"` // the clientID this is set per jodel app version
  DeviceUID string `json:"device_uid"` // the DeviceUID is used to identify the user
  Location Location `json:"location"` // Location struct
}

//voteResponse ...
type voteResponse struct {
  Post SinglePostResponse `json:"post"`
  VoteCount int64 `json:"vote_count"`
}

//ReplyToPost is used to reply to a post
type ReplyToPost struct {
  Ancestor string `json:"ancestor"` // parent post id
  Color string `json:"color"` // HEX color of the post
  Location Location `json:"location"` // Location
  Message string `json:"message"` // the message
}

//ComboPostsResponse ...
type ComboPostsResponse struct {
  Max int64 `json:"max"`
  Recent []SinglePostResponse `json:"recent"` // your most recent posts
  Replied []SinglePostResponse `json:"replied"` // your most replied posts
  Voted []SinglePostResponse `json:"voted"` // your most voted posts
}

//Pagination ...
type Pagination struct {
  Skip int64 `json:"skip"` // aka offset
  Limit int64 `json:"limit"` // limit to x posts
}

//KarmaResponse holds the users Karma
type KarmaResponse struct {
  Karma int64 `json:"karma"` // your karma
}

//NewAccessTokenRequest is used to request a new AccessToken
type NewAccessTokenRequest struct {
  CurrentClientID string `json:"current_client_id"`
  DistinctID string `json:"distinct_id"`
  RefreshToken string `json:"refresh_token"`
}

//NewAccessTokenResonse is used to request a new AccessToken
type NewAccessTokenResonse struct {
  AccessToken string `json:"access_token"`
  ExpirationDate string `json:"expiration_date"`
}
