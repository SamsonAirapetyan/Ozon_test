package postgres

import (
	"Ozon/config"
	"Ozon/internal/utils"
	"Ozon/pkg/logger"
	"context"
	"fmt"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"time"
)

// OpenPoolConnection open pool connection
func OpenPoolConnection(ctx context.Context, cfg *config.Config) (conn *pgxpool.Pool) {
	log := logger.GetLogger()
	err := utils.ConnectionAttemps(func() error {
		var err error
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		conn, err = pgxpool.New(ctx, fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			cfg.PostgresDB.User,
			cfg.PostgresDB.Password,
			cfg.PostgresDB.Host,
			cfg.PostgresDB.Port,
			cfg.PostgresDB.DBName,
			cfg.PostgresDB.SSLmode,
		))
		return err
	}, 3, time.Duration(5)*time.Second)

	if err != nil {
		log.Error("[ERROR] Didn't manage to make connection with database", "message", err.Error())
		os.Exit(1)
	}

	log.Info("[INFO] Database connection is established successfully.")
	return conn
}
