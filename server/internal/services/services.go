package services

import (
	"context"
	"strconv"

	"github.com/arvph/test_tasks/internal/modules"
	"github.com/arvph/test_tasks/internal/repository"
)

// Services представляет репозиторий для работы с базой данных.
type Services struct {
	rp *repository.Repository
}

// NewService создает и возвращает новый экземпляр Services.
func NewService(RP *repository.Repository) *Services {
	return &Services{
		rp: RP,
	}
}

// Create добавляет новую запись в репозиторий.
func (s *Services) Create(ctx context.Context, task modules.Task) error {
	if err := s.rp.Create(ctx, task.UserID, task.Text, task.Status); err != nil {
		return err
	}
	return nil
}

// GetAll получает все записи из репозитория.
func (s *Services) GetAll(ctx context.Context, userID, page, size string) ([]modules.Task, error) {
	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}

	pageSize, err := strconv.Atoi(size)
	if err != nil {
		return nil, err
	}

	tasks, err := s.rp.GetAll(ctx, userID, pageNumber, pageSize)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetByID получает запись из репозитория по ID и userID.
func (s *Services) GetByID(ctx context.Context, userID, ID string) (modules.Task, error) {
	task, err := s.rp.GetByID(ctx, ID, userID)
	if err != nil {
		return task, err
	}
	return task, nil
}

// UpdateByID изменяет запись из репозитория по ID и userID.
func (s *Services) UpdateByID(ctx context.Context, id string, task modules.Task) (modules.Task, error) {
	// ID := strconv.Itoa(id)
	user := strconv.Itoa(task.UserID)

	task, err := s.rp.UpdateByID(ctx, id, user, task.Text, task.Status)
	if err != nil {
		return task, err
	}
	return task, nil
}

// DeleteByID удаляет запись из репозитория по ID и userID.
func (s *Services) DeleteByID(ctx context.Context, userID, ID string) error {
	_, err := s.rp.DeleteByID(ctx, ID, userID)
	if err != nil {
		return err
	}
	return nil
}
