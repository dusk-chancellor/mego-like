package main

import (
	"context"
	"fmt"
	"github.com/dusk-chancellor/mego-like/internal/clients"
	"github.com/jmoiron/sqlx"
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
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.LoadConfig("./dev.env")
	ctx := context.Background()
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Failed to close database connection: %v", err)
		}
	}(db)

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: cfg.RedisPassword,
		DB:       0,
	})

	err = rdb.Ping(ctx).Err()
	if err != nil {
		log.Fatal("redis connection error")
	}
	log.Println("Connected to Redis")

	postClient, err := clients.NewPostClient(cfg.PostServiceAddress)
	if err != nil {
		log.Fatalf("Failed to connect to post service: %v", err)
	}
	commentClient, err := clients.NewCommentClient(cfg.CommentServiceAddress)
	if err != nil {
		log.Fatalf("Failed to connect to comment service: %v", err)
	}

	likeRepo := repositories.NewLikeRepository(db, rdb)
	likeService := services.NewLikeService(likeRepo, postClient, commentClient)

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
