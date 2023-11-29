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
	Get(ctx context.Context, shortUrl string) (string, error)
	Create(ctx context.Context, shortURL, url string) error
}
type App struct {
	handlers   *handlers.Handler
	service    *service.Service
	repository Repository
}

func Run() {
	logger := logger.GetLogger()
	logger.Info("[INFO] Server is starting...")
	cfg := config2.ParseConfig(config2.ConfigViper())
	ctx := context.Background()
	gs := grpc.NewServer()

	connPool := postgres.OpenPoolConnection(ctx, cfg)
	if err := connPool.Ping(ctx); err != nil {
		logger.Info("Unable to ping the database connection", "error", err)
		os.Exit(1)
	}
	//postgres.RunMigrationsUp(ctx, cfg)

	storage := repository.NewStorage(connPool)

	repos := repository.NewPostgresRepository(storage)
	serv := service.NewService(repos)
	hand := handlers.NewHandler(serv)

	protos.RegisterLinkServer(gs, hand)
	reflection.Register(gs)

	l, err := net.Listen("tcp", cfg.Grpc.Address)
	if err != nil {
		logger.Error("[ERROR] Unable to listen", "errors", err)
		os.Exit(1)
	}

	//Like listen and Serve in HTTP
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
