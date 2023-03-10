package v2

import (
	"time"

	"github.com/u8views/go-u8views/internal/models"
)

const (
	appVersion = 1
)

type ProfileView struct {
	ID                   int64
	SocialProviderUserID string
	Username             string
	Name                 string
}

type ProfileViewsStats = models.ProfileViewsStats

type FullProfileView struct {
	ProfileView
	ProfileViewsStats
	CreatedAt time.Time
}

func (v ProfileView) GetName() string {
	if v.Name == "" {
		return v.Username
	}

	return v.Name
}
