package app

import (
	config2 "Ozon/config"
	"Ozon/internal/handlers"
	"Ozon/internal/repository"
	"Ozon/internal/service"
	logger "Ozon/pkg/logger"
	"Ozon/pkg/postgres"
	protos "Ozon/protos/links"
	"context"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Repository interface {
	GetFullLink(ctx context.Context, shortUrl string) (string, error)
	CreateShortLink(ctx context.Context, fullLink, shortLink string) error
}
type App struct {
	handlers   *handlers.Handler
	service    *service.Service
	repository Repository
}

func Run(a *App) {
	logger := logger.GetLogger()
	logger.Info("[INFO] Server is starting...")
	cfg := config2.ParseConfig(config2.ConfigViper())
	ctx := context.Background()
	gs := grpc.NewServer()

	//load .env file
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Can not load .env file")
	}

	stor := os.Getenv("STORAGE_TYPE")

	app := &App{}
	switch stor {
	case "psql":
		connPool := postgres.OpenPoolConnection(ctx, cfg)
		if err := connPool.Ping(ctx); err != nil {
			logger.Info("Unable to ping the database connection", "error", err)
			os.Exit(1)
		}
		//postgres.RunMigrationsUp(ctx, cfg)

		storage := repository.NewStorage(connPool)
		app.repository = repository.NewPostgresRepository(storage)
		logger.Info("Postgres storage")
	case "inMemo":
		app.repository = repository.NewMemoryRepository()
		logger.Info("In-memory storage")
	default:
		logger.Error("No database has chosen")
	}
	app.service = service.NewService(a.repository)
	app.handlers = handlers.NewHandler(a.service)

	protos.RegisterLinkServer(gs, a.handlers)
	reflection.Register(gs)

	l, err := net.Listen("tcp", cfg.Grpc.Address)
	if err != nil {
		logger.Error("[ERROR] Unable to listen", "errors", err)
		os.Exit(1)
	}

	gs.Serve(l)
}

func StartServer(ctx context.Context, srv http.Server) {
	logger := logger.GetLogger()

	go func() {
		logger.Info("Starting server...")
		err := srv.ListenAndServe()
		if err != nil {
			logger.Error("Server was stopped", "error", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	signal := <-sigChan
	logger.Info("signal has been recieved", "signal", signal)
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	srv.Shutdown(ctx)
}
