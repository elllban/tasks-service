package main

import (
	"log"

	"github.com/elllban/tasks-service/internal/database"
	"github.com/elllban/tasks-service/internal/task"
	"github.com/elllban/tasks-service/internal/transport/grpc"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := task.NewRepository(db)
	svc := task.NewService(repo)

	userClient, conn, err := grpc.NewUserClient("localhost:50051")
	if err != nil {
		log.Fatalf("failed to connect to users: %v", err)
	}
	defer conn.Close()

	if err = grpc.RunGRPC(svc, userClient); err != nil {
		log.Fatalf("Tasks gRPC server error: %v", err)
	}
}
