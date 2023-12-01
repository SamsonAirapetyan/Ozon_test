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
	log := logger.GetLogger()
	err := godotenv.Load(".env")
	if err != nil {
		log.Error("[ERROR] Can not load .env file")
	}
	storage := os.Getenv("STORAGE_TYPE")
	switch storage {
	case "psql":
		connPool := postgres.OpenPoolConnection(ctx, cfg)
		if err := connPool.Ping(ctx); err != nil {
			log.Error("[ERROR] Unable to ping the database connection", "error", err)
			os.Exit(1)
		}
		stor := repository.NewStorage(connPool)
		app.repository = repository.NewPostgresRepository(stor)
		log.Info("[INFO] Postgres storage")
	case "inMemory":
		app.repository = repository.NewMemoryRepository()
		log.Info("[INFO] In-memory storage")
	default:
		log.Error("[ERROR] No database has chosen")
	}
	app.service = service.NewService(app.repository)
}

func Run() {
	log := logger.GetLogger()
	log.Info("[INFO] Server is starting...")
	ctx := context.Background()
	cfg := config2.ParseConfig(config2.ConfigViper())

	app := &App{}
	New(ctx, cfg, app)
	app.handlers = handlers.NewHandler(app.service)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	lis, err := net.Listen(cfg.Grpc.Network, cfg.Grpc.Address)
	if err != nil {
		log.Info("[ERROR] Failed to listen")
	}

	s := grpc.NewServer()
	protos.RegisterLinkServer(s, app.handlers)
	go func() {
		if err = s.Serve(lis); err != nil {
			log.Error("[ERROR] failed to serve: " + err.Error())
		}
	}()

	log.Error("[ERROR] Failed to dial server: " + err.Error())
	conn, err := grpc.DialContext(
		context.Background(),
		cfg.Grpc.Address,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
	}
	defer conn.Close()

	runRest(ctx, conn, interrupt)

	select {
	case <-interrupt:
		log.Info("[INFO] Signal has been received", "signal", interrupt)
		s.GracefulStop()
		return
	case <-ctx.Done():
		return
	}
}

func runRest(ctx context.Context, conn *grpc.ClientConn, shutdown chan os.Signal) {
	log := logger.GetLogger()
	mux := runtime.NewServeMux()
	err := protos.RegisterLinkHandler(context.Background(), mux, conn)
	if err != nil {
		log.Error("[ERROR] Failed to register gateway:" + err.Error())
	}

	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Info("[INFO] Serving gRPC-Gateway on port " + ":8080")
	go func() {
		if err = gwServer.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Error("[ERROR] Server closed: " + err.Error())
			} else {
				log.Error("[ERROR] Failed to listen and serve: " + err.Error())
			}
			close(shutdown)
		}
	}()
	select {
	case <-shutdown:
		log.Info("[INFO] Reeived interrupt stgnal. Shutting down gRPC-Gateway...")
		if err := gwServer.Shutdown(context.Background()); err != nil {
			log.Error("[ERROR] Error during shutdown:", err)
		}
		return
	case <-ctx.Done():
		return
	}
}
