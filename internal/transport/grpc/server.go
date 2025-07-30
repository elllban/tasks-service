package grpc

import (
	"net"

	taskpb "github.com/elllban/project-protos/proto/task"
	userpb "github.com/elllban/project-protos/proto/user"
	"github.com/elllban/tasks-service/internal/task"
	"google.golang.org/grpc"
)

func RunGRPC(svc task.Service, uc userpb.UserServiceClient) error {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	taskpb.RegisterTaskServiceServer(grpcServer, NewHandler(svc, uc))

	return grpcServer.Serve(lis)
}
