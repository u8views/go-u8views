package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/u8views/go-u8views/internal/models/oauth2"
	"github.com/u8views/go-u8views/internal/models/oauth2/github"

	"github.com/valyala/fasthttp"
)

var (
	defaultClient = &fasthttp.Client{
		ReadBufferSize: 16 * 1024,
	}
)

func User(secret oauth2.Secret, code string) (result oauth2.SocialProviderUser, err error) {
	accessTokenResponse, err := getAccessToken(secret, code)
	if err != nil {
		return result, err
	}

	userResponse, err := getUser(accessTokenResponse.AccessToken)
	if err != nil {
		return result, err
	}

	return oauth2.SocialProviderUser{
		ID:       strconv.FormatUint(userResponse.ID, 10),
		Username: userResponse.Login,
		Name:     userResponse.Name,
	}, nil
}

// https://developer.github.com/apps/building-oauth-apps/authorizing-oauth-apps/
func Redirect(secret oauth2.Secret) string {
	var uri = fasthttp.AcquireURI()

	uri.Update(github.Endpoint.AuthURL)
	var args = uri.QueryArgs()
	args.Set("client_id", secret.ClientID)
	// redirect_uri already set on https://github.com/organizations/u8views/settings/applications
	// args.Set("redirect_uri", "https://u8views.com/oauth/callback/github")
	args.Set("scope", secret.Scope)

	var result = uri.String()

	fasthttp.ReleaseURI(uri)

	return result
}

func getAccessToken(secret oauth2.Secret, code string) (result github.AccessTokenResponse, err error) {
	var (
		request  = fasthttp.AcquireRequest()
		response = fasthttp.AcquireResponse()
	)

	defer func() {
		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)
	}()

	request.Header.SetMethod("POST")
	request.Header.Set("accept", "application/json")
	request.SetRequestURI(github.Endpoint.TokenURL)

	q := request.URI().QueryArgs()
	q.Add("client_id", secret.ClientID)
	q.Add("client_secret", secret.ClientSecret)
	q.Add("code", code)

	err = defaultClient.Do(request, response)
	if err != nil {
		return result, fmt.Errorf("GitHub (get access token) http err %v", err)
	}

	if response.StatusCode() != http.StatusOK {
		return result, fmt.Errorf("GitHub (get access token) unexpected http status code %d", response.StatusCode())
	}

	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return result, fmt.Errorf("GitHub (get access token) cannot unmarshal response err: %v", err)
	}

	return result, nil
}

func getUser(accessToken string) (result github.User, err error) {
	var (
		request  = fasthttp.AcquireRequest()
		response = fasthttp.AcquireResponse()
	)

	defer func() {
		fasthttp.ReleaseRequest(request)
		fasthttp.ReleaseResponse(response)
	}()

	request.SetRequestURI(github.ApiUserURI)
	request.Header.Set("Authorization", "token "+accessToken)

	err = defaultClient.Do(request, response)
	if err != nil {
		return result, fmt.Errorf("GitHub (get user) http err %v", err)
	}

	if response.StatusCode() != http.StatusOK {
		return result, fmt.Errorf("GitHub (get user) unexpected http status code %d", response.StatusCode())
	}

	err = json.Unmarshal(response.Body(), &result)
	if err != nil {
		return result, fmt.Errorf("GitHub (get user) cannot unmarshal response err: %v", err)
	}

	return result, nil
}
