package services

import (
	"context"
	"database/sql"
	"log"
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

func (s *UserService) GetByID(ctx context.Context, id int64) (dbs.UsersGetByIDRow, error) {
	return s.repository.Queries().UsersGetByID(ctx, id)
}

func (s *UserService) Users(ctx context.Context, limit int32) ([]dbs.UsersGetRow, error) {
	return s.repository.Queries().UsersGet(ctx, limit)
}

func (s *UserService) GetAllUsernames(ctx context.Context) ([]dbs.UsersGetRow, error) {
	return s.repository.Queries().UsersGet(ctx, 100)
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
	var referral bool

	if referrerUserID > 0 {
		_, err := s.repository.Queries().UsersGetBySocialProvider(ctx, dbs.UsersGetBySocialProviderParams{
			SocialProvider:       provider,
			SocialProviderUserID: socialProviderUserID,
		})
		if err == sql.ErrNoRows {
			// NEW

			referral = true
		} else if err != nil {
			log.Printf("Database err %s\n", err)
		} else {
			// ALREADY EXISTS
		}
	}

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

		if referral {
			err := queries.ReferralsNew(ctx, dbs.ReferralsNewParams{
				RefereeUserID:  txID,
				ReferrerUserID: referrerUserID,
			})
			if err != nil {
				log.Printf("Database err %s\n", err)

				// NOP
			}
		}

		id = txID

		return nil
	})
	if err != nil {
		return id, err
	}

	return id, nil
}
