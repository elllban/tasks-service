package grpc

import (
	"context"
	"fmt"

	taskpb "github.com/elllban/project-protos/proto/task"
	userpb "github.com/elllban/project-protos/proto/user"
	"github.com/elllban/tasks-service/internal/task"
)

type Handler struct {
	svc        task.Service
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

func NewHandler(svc task.Service, uc userpb.UserServiceClient) *Handler {
	return &Handler{svc: svc, userClient: uc}
}

func (h *Handler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserId, err)
	}

	taskRequest := task.TaskRequest{
		UserID: req.UserId,
		Task:   req.Task,
		IsDone: req.IsDone,
	}

	taskResponse, err := h.svc.CreateTask(taskRequest)
	if err != nil {

		return nil, err
	}

	response := &taskpb.CreateTaskResponse{Task: &taskpb.Task{
		Id:     taskResponse.ID,
		Task:   taskResponse.Task,
		IsDone: taskResponse.IsDone,
		UserId: taskResponse.UserID,
	}}

	return response, nil
}

func (h *Handler) GetTask(_ context.Context, req *taskpb.GetTaskRequest) (*taskpb.GetTaskResponse, error) {
	taskResponse, err := h.svc.GetTaskByID(req.Id)
	if err != nil {
		return &taskpb.GetTaskResponse{}, err
	}

	response := &taskpb.GetTaskResponse{Task: &taskpb.Task{
		Id:     taskResponse.ID,
		Task:   taskResponse.Task,
		IsDone: taskResponse.IsDone,
		UserId: taskResponse.UserID,
	}}

	return response, nil
}

func (h *Handler) ListTasks(_ context.Context, _ *taskpb.ListTasksRequest) (*taskpb.ListTasksResponse, error) {
	tasks, err := h.svc.GetAllTasks()
	if err != nil {
		return &taskpb.ListTasksResponse{}, err
	}

	response := &taskpb.ListTasksResponse{}

	for _, us := range tasks {
		tk := &taskpb.Task{
			Id:     us.ID,
			Task:   us.Task,
			IsDone: us.IsDone,
			UserId: us.UserID,
		}
		response.Tasks = append(response.Tasks, tk)
	}

	return response, nil
}

func (h *Handler) ListTasksByUser(_ context.Context, req *taskpb.ListTasksByUserRequest) (*taskpb.ListTasksByUserResponse, error) {
	tasks, err := h.svc.GetTasksByUser(req.UserId)
	if err != nil {
		return &taskpb.ListTasksByUserResponse{}, err
	}

	response := &taskpb.ListTasksByUserResponse{}

	for _, us := range tasks {
		tk := &taskpb.Task{
			Id:     us.ID,
			Task:   us.Task,
			IsDone: us.IsDone,
			UserId: us.UserID,
		}
		response.Tasks = append(response.Tasks, tk)
	}

	return response, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserId, err)
	}

	taskRequest := task.TaskRequest{
		Task:   req.Task,
		IsDone: req.IsDone,
		UserID: req.UserId,
	}

	taskResponse, err := h.svc.UpdateTask(req.Id, taskRequest)
	if err != nil {
		return &taskpb.UpdateTaskResponse{}, err
	}

	response := &taskpb.UpdateTaskResponse{Task: &taskpb.Task{
		Id:     taskResponse.ID,
		Task:   taskResponse.Task,
		IsDone: taskResponse.IsDone,
		UserId: taskResponse.UserID,
	}}

	return response, nil
}

func (h *Handler) DeleteTask(_ context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	if err := h.svc.DeleteTask(req.Id); err != nil {
		return &taskpb.DeleteTaskResponse{}, err
	}

	return &taskpb.DeleteTaskResponse{}, nil
}
