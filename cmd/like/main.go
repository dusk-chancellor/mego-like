package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	adapter "github.com/dusk-chancellor/mego-like/internal/adapters/grpc"
	"github.com/dusk-chancellor/mego-like/internal/config"
	"github.com/dusk-chancellor/mego-like/internal/database"
	"github.com/dusk-chancellor/mego-like/internal/repositories"
	"github.com/dusk-chancellor/mego-like/internal/services"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig()
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	likeRepo := repositories.NewLikeRepository(db)
	likeLocalCache := services.NewLikeLocalCache()
	likeService := services.NewLikeService(likeRepo, likeLocalCache)

	l, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	adapter.RegisterGRPC(grpcServer, likeService, cfg)
	go log.Fatal(grpcServer.Serve(l))

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop
	grpcServer.GracefulStop()
}
