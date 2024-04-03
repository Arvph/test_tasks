package repository

import (
	"context"

	"github.com/arvph/test_tasks/internal/database"
	"github.com/arvph/test_tasks/internal/modules"
)

// Repository представляет репозиторий для работы с базой данных.
type Repository struct {
	db *database.DB
}

// NewRepository создает и возвращает новый экземпляр Repository.
func NewRepository(DB *database.DB) *Repository {
	return &Repository{
		db: DB,
	}
}

// Create добавляет новую запись в репозиторий.
func (rp *Repository) Create(ctx context.Context, userID int, text, status string) error {
	sql := `INSERT INTO "task" (user_id, text, created_at, status) VALUES ($1, $2, CURRENT_TIMESTAMP, $3);`

	_, err := rp.db.Pool.Exec(ctx, sql, userID, text, status)
	if err != nil {
		return err
	}
	return nil
}

// GetAll получает все записи из репозитория.
func (rp *Repository) GetAll(ctx context.Context, userID string, pageNumber, pageSize int) ([]modules.Task, error) {
	// Расчет смещения на основе номера страницы и размера страницы
	offset := (pageNumber - 1) * pageSize

	sql := `SELECT id, user_id, text, created_at, status FROM tasks WHERE user_id = $1 LIMIT $2 OFFSET $3;`

	rows, err := rp.db.Pool.Query(ctx, sql, userID, pageSize, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []modules.Task

	for rows.Next() {
		var task modules.Task
		if err := rows.Scan(&task.ID, &task.UserID, &task.Text, &task.CreatedAt, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// GetByID получает запись из репозитория по ID и userID.
func (rp *Repository) GetByID(ctx context.Context, ID, userID string) (modules.Task, error) {
	sql := `SELECT id, user_id, text, created_at, status FROM tasks WHERE id=$1 AND user_id = $2;`

	row := rp.db.Pool.QueryRow(ctx, sql, ID, userID)

	var task modules.Task

	if err := row.Scan(&task.ID, &task.UserID, &task.Text, &task.CreatedAt, &task.Status); err != nil {
		// if err == sql.ErrNoRows {
		// 	return task, fmt.Errorf("no task found with ID %s and userID %s", ID, userID)
		// }
		return task, err
	}

	return task, nil
}

// UpdateByID измен запись из репозитория по ID и userID.
func (rp *Repository) UpdateByID(ctx context.Context, ID, userID, newText, Status string) (modules.Task, error) {
	sql := `UPDATE tasks SET text = $1, status = $2 WHERE id=$3 AND user_id = $4 RETURNING id, user_id, text, created_at, status;`

	row := rp.db.Pool.QueryRow(ctx, sql, newText, Status, ID, userID)

	var task modules.Task

	if err := row.Scan(&task.ID, &task.UserID, &task.Text, &task.CreatedAt, &task.Status); err != nil {
		return task, err
	}

	return task, nil
}

// DeleteByID удаляет запись из репозитория по ID и userID.
func (rp *Repository) DeleteByID(ctx context.Context, ID, userID string) (modules.Task, error) {
	task, err := rp.GetByID(ctx, ID, userID)
	if err != nil {
		return task, err
	}

	sql := `DELETE FROM tasks WHERE id=$1 AND user_id = $2;`
	if _, err := rp.db.Pool.Exec(ctx, sql, ID, userID); err != nil {
		return task, err
	}

	return task, nil
}
