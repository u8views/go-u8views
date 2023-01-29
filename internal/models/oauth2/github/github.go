package github

import "github.com/u8views/go-u8views/internal/models/oauth2"

const (
	ApiUserURI = "https://api.github.com/user"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://github.com/login/oauth/authorize",
	TokenURL: "https://github.com/login/oauth/access_token",
}
