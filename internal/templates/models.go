package templates

import "github.com/u8views/go-u8views/internal/models"

type CurrentUserView struct {
	ID                   int64
	SocialProviderUserID string
	Username             string
	Name                 string
}

type ProfileView struct {
	ID                   int64
	SocialProviderUserID string
	Username             string
	Name                 string
}

type ProfileViewsStats = models.ProfileViewsStats

func (v CurrentUserView) GetName() string {
	if v.Name == "" {
		return v.Username
	}

	return v.Name
}

func (v ProfileView) GetName() string {
	if v.Name == "" {
		return v.Username
	}

	return v.Name
}
