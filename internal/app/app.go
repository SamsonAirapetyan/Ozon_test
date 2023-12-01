package app

import (
	config2 "Ozon/config"
	"Ozon/internal/handlers"
	"Ozon/internal/repository"
	"Ozon/internal/service"
	"Ozon/pkg/logger"
	"Ozon/pkg/postgres"
	protos "Ozon/protos/links"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//go:generate mockgen -source=app.go -destination=../internal/service/mocks/mock.go

type Repository interface {
	GetFullLink(ctx context.Context, shortUrl string) (string, error)
	CreateShortLink(ctx context.Context, fullLink, shortLink string) error
}
type App struct {
	handlers   *handlers.Handler
	service    *service.Service
	repository Repository
}

func New(ctx context.Context, cfg *config2.Config, app *App) {
	godotenv.Load(".env")
	//if err != nil {
	//	logger.Error("Can not load .env file")
	//}
	logger := logger.GetLogger()
	stor := os.Getenv("STORAGE_TYPE")
	switch stor {
	case "psql":
		connPool := postgres.OpenPoolConnection(ctx, cfg)
		if err := connPool.Ping(ctx); err != nil {
			logger.Info("Unable to ping the database connection", "error", err)
			os.Exit(1)
		}
		//postgres.RunMigrationsUp(ctx, cfg)
		//defer postgres.DownMigrationsUp(ctx, cfg)

		storage := repository.NewStorage(connPool)
		app.repository = repository.NewPostgresRepository(storage)
		logger.Info("Postgres storage")
	case "inMemory":
		app.repository = repository.NewMemoryRepository()
		logger.Info("In-memory storage")
	default:
		logger.Error("No database has chosen")
	}
	app.service = service.NewService(app.repository)
}

func Run() {
	logger := logger.GetLogger()
	logger.Info("[INFO] Server is starting...")
	ctx := context.Background()
	cfg := config2.ParseConfig(config2.ConfigViper())

	//load .env file
	//err := godotenv.Load(".env")
	//if err != nil {
	//	logger.Error("Can not load .env file")
	//}

	app := &App{}

	New(ctx, cfg, app)

	postgres.RunMigrationsUp(ctx, cfg)
	defer postgres.DownMigrationsUp(ctx, cfg)
	//stor := os.Getenv("STORAGE_TYPE")
	//
	//switch stor {
	//case "psql":
	//	connPool := postgres.OpenPoolConnection(ctx, cfg)
	//	if err := connPool.Ping(ctx); err != nil {
	//		logger.Info("Unable to ping the database connection", "error", err)
	//		os.Exit(1)
	//	}
	//
	//	postgres.RunMigrationsUp(ctx, cfg)
	//	defer postgres.DownMigrationsUp(ctx, cfg)
	//
	//	storage := repository.NewStorage(connPool)
	//	app.repository = repository.NewPostgresRepository(storage)
	//	logger.Info("Postgres storage")
	//case "inMemory":
	//	app.repository = repository.NewMemoryRepository()
	//	logger.Info("In-memory storage")
	//default:
	//	logger.Error("No database has chosen")
	//}
	//app.service = service.NewService(app.repository)
	app.handlers = handlers.NewHandler(app.service)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	lis, err := net.Listen("tcp", ":9092")
	if err != nil {
		logger.Info("Failed to listen")
	}

	s := grpc.NewServer()
	protos.RegisterLinkServer(s, app.handlers)
	go func() {
		if err = s.Serve(lis); err != nil {
			logger.Error("failed to serve: " + err.Error())
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		":9092",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		logger.Error("Failed to dial server: " + err.Error())
	}
	defer conn.Close()

	runRest(ctx, conn, interrupt)

	select {
	case <-interrupt:
		logger.Info("signal has been recieved", "signal", interrupt)
		s.GracefulStop()
		return
	case <-ctx.Done():
		return
	}
}

func runRest(ctx context.Context, conn *grpc.ClientConn, shutdown chan os.Signal) {
	logger := logger.GetLogger()
	gwmux := runtime.NewServeMux()
	err := protos.RegisterLinkHandler(context.Background(), gwmux, conn)
	if err != nil {
		logger.Error("Failed to register gateway:" + err.Error())
	}

	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}

	logger.Info("Serving gRPC-Gateway on port " + ":8080")
	go func() {
		if err = gwServer.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				logger.Error("Server closed: " + err.Error())
			} else {
				logger.Error("Failed to listen and serve: " + err.Error())
			}
			close(shutdown)
		}
	}()
	select {
	case <-shutdown:
		logger.Info("Reeived interrupt stgnal. Shutting down gRPC-Gateway...")
		if err := gwServer.Shutdown(context.Background()); err != nil {
			logger.Error("Error during shutdown:", err)
		}
		return
	case <-ctx.Done():
		return
	}
}
