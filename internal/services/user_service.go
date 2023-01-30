package services

import (
	"context"
	"time"

	"github.com/u8views/go-u8views/internal/db"
	"github.com/u8views/go-u8views/internal/storage/dbs"
)

type UserService struct {
	repository *db.Repository
}

func NewUserService(repository *db.Repository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) GetBySocialProvider(ctx context.Context, provider dbs.SocialProvider, socialProviderUserID string) (int64, error) {
	return s.repository.Queries().UsersGetBySocialProvider(ctx, dbs.UsersGetBySocialProviderParams{
		SocialProvider:       provider,
		SocialProviderUserID: socialProviderUserID,
	})
}

func (s *UserService) GetBySocialProviderUsername(ctx context.Context, provider dbs.SocialProvider, username string) (dbs.UsersGetBySocialProviderUsernameRow, error) {
	return s.repository.Queries().UsersGetBySocialProviderUsername(ctx, dbs.UsersGetBySocialProviderUsernameParams{
		SocialProvider: provider,
		Username:       username,
	})
}

func (s *UserService) GetByID(ctx context.Context, id int64) (dbs.UsersGetByIDRow, error) {
	return s.repository.Queries().UsersGetByID(ctx, id)
}

func (s *UserService) Users(ctx context.Context, limit int32) ([]dbs.UsersGetRow, error) {
	return s.repository.Queries().UsersGet(ctx, limit)
}

func (s *UserService) Upsert(
	ctx context.Context,
	provider dbs.SocialProvider,
	socialProviderUserID string,
	username string,
	name string,
) (id int64, err error) {
	now := time.Now().UTC()

	err = s.repository.WithTransaction(ctx, func(queries *dbs.Queries) error {
		txID, err := queries.UsersNew(ctx, dbs.UsersNewParams{
			SocialProvider:       provider,
			SocialProviderUserID: socialProviderUserID,
			Username:             username,
			Name:                 name,
			CreatedAt:            now,
			UpdatedAt:            now,
			LastLoginAt:          now,
		})
		if err != nil {
			return err
		}

		err = queries.UsersUpdateUsername(ctx, dbs.UsersUpdateUsernameParams{
			Username:  username,
			Name:      name,
			UpdatedAt: now,
			ID:        txID,
		})
		if err != nil {
			return err
		}

		err = queries.ProfileTotalViewsNew(ctx, txID)
		if err != nil {
			return err
		}

		id = txID

		return nil
	})
	if err != nil {
		return id, err
	}

	return id, nil
}
