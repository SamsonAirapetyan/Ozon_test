package repository

import (
	errors "Ozon/domain"
	"Ozon/pkg/logger"
	"context"
	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v5"
	"time"
)

type PostgresRepository struct {
	logger  hclog.Logger
	storage *Storage //poolConnection
}

func NewPostgresRepository(storage *Storage) *PostgresRepository {
	return &PostgresRepository{logger: logger.GetLogger(), storage: storage}
}

func (r *PostgresRepository) GetFullLink(ctx context.Context, shortLink string) (string, error) {
	r.logger.Info("[INFO] Get Links from Postgres DataBase")
	query := `SELECT * FROM link WHERE short_link = $1 `
	entity := &Link{}

	conn, err := r.storage.GetPgConnPool().Acquire(ctx)
	if err != nil {
		return "", err
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return "", err
	}
	defer tx.Rollback(ctx)

	if err = tx.QueryRow(ctx, query, shortLink).Scan(&entity.Full_link, &entity.Short_link); err != nil {
		if err == pgx.ErrNoRows {
			r.logger.Error("[ERROR NOT FOUND]  detected")
			return "", errors.ErrNoRecordFound
		}
		return "", err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return "", err
	}
	return entity.Full_link, nil
}

func (r *PostgresRepository) CreateShortLink(ctx context.Context, fullLink, shortLink string) error {
	r.logger.Info("[INFO] Set Links in Postgres DataBase")
	query := `INSERT INTO link(full_link,short_link) VALUES ($1,$2) ON CONFLICT DO NOTHING`
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	conn, er := r.storage.GetPgConnPool().Acquire(ctx)
	if er != nil {
		return er
	}
	defer conn.Release()
	tx, er := conn.Begin(ctx)
	if er != nil {
		return er
	}
	defer tx.Rollback(ctx)
	_, err := tx.Exec(ctx, query, fullLink, shortLink)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}
