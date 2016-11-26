package godel

//TODO: restructure this file and rename some structs to be more clear ex: LocationCoordinates => Coordinates etc..

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
  Location LocationResponse `json:"location"` //Location struct
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
  ChildCount int64 `json:"child_count"`
  Children []SingleComment `json:"children"`
  Color string `json:"color"` // HEX Color of the post
  CreatedAt string `json:"created_at"` // ISO Timestamp
  Discovered int64 `json:"discovered"` // unknown
  DiscoveredBy int64 `json:"discovered_by"` // unknown
  Distance int64 `json:"distance"` // distance to current location
  GotThanks bool `json:"got_thanks"` // bool
  Location LocationResponse `json:"location"`
  Message string `json:"message"`
  NotificationsEnabled bool `json:"notifications_enabled"`
  PinCount int64 `json:"pin_count"`
  PostID string `json:"post_id"`
  PostOwn string `json:"post_own"`
  // Tags
  UpdatedAt string `json:"updated_at"`
  UserHandle string `json:"user_handle"`
  VoteCount int64 `json:"vote_count"`
}

// SendLocationRequest is used to set a new Location for a user
type SendLocationRequest struct {
  Location LocationResponse `json:"location"`
}

// NewPost can be used to create a new Post
type NewPost struct {
  Color string `json:"color"`
  Location LocationResponse `json:"location"`
  Message string `json:"message"`
}

//LocationCoordinates hold Latitude and Longitude of the current location
type LocationCoordinates struct {
  Lat float64 `json:"lat"`
  Lng float64 `json:"lng"`
}

//LocationResponse Holds the users current Location
type LocationResponse struct {
  LocAccuracy float64 `json:"loc_accuracy"`
  City string `json:"city"`
  Country string `json:"country"`
  Name string `json:"name"`
  LocCoordinates LocationCoordinates `json:"loc_coordinates"`
}

//LoginPayload is used to log the user in
type LoginPayload struct {
  ClientID string `json:"client_id"`
  DeviceUID string `json:"device_uid"`
  Location LocationResponse `json:"location"`
}

//VoteResponse ...
type VoteResponse struct {
  Post SinglePostResponse `json:"name"`
  VoteCount int64 `json:"vote_count"`
}

//VoteResponse ...
type ReplyToPost struct {
  Ancestor string `json:"ancestor"`
  Color string `json:"color"`
  Location LocationResponse `json:"location"`
  Message string `json:"message"`
}

//ComboPosts ...
type ComboPosts struct {
  Max int64 `json:"max"`
  Recent []SinglePostResponse `json:"recent"`
  Replied []SinglePostResponse `json:"replied"`
  Voted []SinglePostResponse `json:"voted"`
}

//Pagination ...
type Pagination struct {
  Skip int64 `json:"skip"`
  Limit int64 `json:"limit"`
}

//Karma holds the users Karma
type KarmaResponse struct {
  Karma int64 `json:"karma"`
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
