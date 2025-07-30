package task

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateTask(res TaskResponse) error
	GetAllTasks() ([]TaskResponse, error)
	GetTasksByUser(userId string) ([]TaskResponse, error)
	GetTaskByID(id string) (TaskResponse, error)
	UpdateTask(res TaskResponse) error
	DeleteTask(id string) error
}

type tsRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &tsRepository{db: db}
}

func (r *tsRepository) CreateTask(res TaskResponse) error {
	return r.db.Create(&res).Error
}

func (r *tsRepository) GetAllTasks() ([]TaskResponse, error) {
	var tasks []TaskResponse
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *tsRepository) GetTasksByUser(userId string) ([]TaskResponse, error) {
	var tasks []TaskResponse
	err := r.db.Find(&tasks, "user_id = ?", userId).Error
	return tasks, err
}

func (r *tsRepository) GetTaskByID(id string) (TaskResponse, error) {
	var ts TaskResponse
	err := r.db.First(&ts, "id = ?", id).Error
	return ts, err
}

func (r *tsRepository) UpdateTask(res TaskResponse) error {
	return r.db.Save(&res).Error
}

func (r *tsRepository) DeleteTask(id string) error {
	return r.db.Delete(&TaskResponse{}, "id = ?", id).Error
}
