package task

import "github.com/google/uuid"

type Service interface {
	CreateTask(req TaskRequest) (TaskResponse, error)
	GetAllTasks() ([]TaskResponse, error)
	GetTasksByUser(userId string) ([]TaskResponse, error)
	GetTaskByID(id string) (TaskResponse, error)
	UpdateTask(id string, req TaskRequest) (TaskResponse, error)
	DeleteTask(id string) error
}

type tsService struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &tsService{repo: r}
}

func (s *tsService) CreateTask(req TaskRequest) (TaskResponse, error) {
	newTask := TaskResponse{
		ID:     uuid.NewString(),
		Task:   req.Task,
		IsDone: req.IsDone,
		UserID: req.UserID,
	}

	if err := s.repo.CreateTask(newTask); err != nil {
		return TaskResponse{}, err
	}

	return newTask, nil
}

func (s *tsService) GetAllTasks() ([]TaskResponse, error) {
	tasks, err := s.repo.GetAllTasks()
	if err != nil {
		return []TaskResponse{}, err
	}

	return tasks, nil
}

func (s *tsService) GetTasksByUser(userId string) ([]TaskResponse, error) {
	tasks, err := s.repo.GetTasksByUser(userId)
	if err != nil {
		return []TaskResponse{}, err
	}

	return tasks, nil
}

func (s *tsService) GetTaskByID(id string) (TaskResponse, error) {
	newTask, err := s.repo.GetTaskByID(id)
	if err != nil {
		return TaskResponse{}, err
	}

	return newTask, nil
}

func (s *tsService) UpdateTask(id string, req TaskRequest) (TaskResponse, error) {
	newTasks, err := s.repo.GetTaskByID(id)
	if err != nil {
		return TaskResponse{}, err
	}

	newTasks.Task = req.Task
	newTasks.IsDone = req.IsDone
	newTasks.UserID = req.UserID

	if err = s.repo.UpdateTask(newTasks); err != nil {
		return TaskResponse{}, err
	}

	return newTasks, nil
}

func (s *tsService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
