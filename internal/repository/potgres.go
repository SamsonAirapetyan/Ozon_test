package repository

import (
	"Ozon/pkg/logger"
	"context"
	"github.com/hashicorp/go-hclog"
	"time"
)

type PostgresRepository struct {
	logger  hclog.Logger
	storage *Storage //poolConnection
}

// убрать от сюда
type Link struct {
	Full_link  string `db:"id"`
	Short_link string `db:"privilege_title"`
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
	if err := r.checklink(ctx, fullLink); err != nil {
		query := `INSERT INTO link(full_link,short_link) VALUES ($1,$2)`
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
		_, err = tx.Exec(ctx, query, fullLink, shortLink)
		if err != nil {
			return err
		}

		err = tx.Commit(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *PostgresRepository) checklink(ctx context.Context, fullLink string) error {
	query := `SELECT * FROM link WHERE full_link = $1`

	entity := &Link{}
	conn, err := r.storage.GetPgConnPool().Acquire(ctx)
	if err != nil {
		return err
	}
	defer conn.Release()
	tx, err := conn.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if err = tx.QueryRow(ctx, query, fullLink).Scan(&entity.Full_link, &entity.Short_link); err != nil {
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}
