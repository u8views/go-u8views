package services

import (
	"context"
	"time"

	"github.com/u8views/go-u8views/internal/db"
	"github.com/u8views/go-u8views/internal/models"
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

func (s *UserService) GetUsernameByOldUsername(ctx context.Context, provider dbs.SocialProvider, oldUsername string) (string, error) {
	return s.repository.Queries().UsernameHistoryGetByOldUsername(ctx, dbs.UsernameHistoryGetByOldUsernameParams{
		Username:       oldUsername,
		SocialProvider: provider,
	})
}

func (s *UserService) GetByID(ctx context.Context, id int64) (dbs.UsersGetByIDRow, error) {
	return s.repository.Queries().UsersGetByID(ctx, id)
}

func (s *UserService) Users(ctx context.Context, limit int32) ([]dbs.UsersGetRow, error) {
	return s.repository.Queries().UsersGet(ctx, limit)
}

func (s *UserService) UsersCount(ctx context.Context) (int64, error) {
	return s.repository.Queries().UsersCount(ctx)
}

func (s *UserService) Usernames(ctx context.Context, offset, limit int32) ([]string, error) {
	return s.repository.Queries().Usernames(ctx, dbs.UsernamesParams{
		Offset: offset,
		Limit:  limit,
	})
}

func (s *UserService) UsersCreatedAtStatsByDay(ctx context.Context) ([]models.TimeCount, error) {
	to := time.Now().UTC().Truncate(24 * time.Hour)
	from := to.AddDate(0, -1, 0)

	rows, err := s.repository.Queries().UsersCreatedAtStatsByDay(ctx, dbs.UsersCreatedAtStatsByDayParams{
		From: from,
		To:   to,
	})
	if err != nil {
		return nil, err
	}

	result := make([]models.TimeCount, len(rows))
	for i, row := range rows {
		result[i] = models.TimeCount{
			Time:  row.Time.Unix(),
			Count: row.Count,
		}
	}
	return result, nil
}

func (s *UserService) Upsert(
	ctx context.Context,
	provider dbs.SocialProvider,
	socialProviderUserID string,
	username string,
	name string,
	referrerUserID int64,
) (id int64, err error) {
	row, err := s.repository.Queries().UsersUpsert(ctx, dbs.UsersUpsertParams{
		SocialProvider:       provider,
		SocialProviderUserID: socialProviderUserID,
		Username:             username,
		Name:                 name,
		Now:                  time.Now().UTC(),
		ReferrerUserID:       referrerUserID,
	})
	if err != nil {
		return 0, err
	}

	return row.ID, nil
}
