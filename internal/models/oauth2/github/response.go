package github

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type User struct {
	ID    uint64 `json:"id"`
	Login string `json:"login"`
}
