package repository

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"testCase"
)

type TaskPostgres struct {
	db *sqlx.DB
}

var (
	ErrUserNotFound = errors.New("user not found")
	ErrTaskNotFound = errors.New("task not found")
)

func NewTaskPostgres(db *sqlx.DB) *TaskPostgres {
	return &TaskPostgres{db: db}
}
func (r *TaskPostgres) GetCompletedTasksByUserID(userID string) ([]testCase.Task, error) {
	var tasks []testCase.Task

	query := `
		SELECT t.id, t.name, t.description, t.reward_points
		FROM tasks t
		JOIN user_tasks ut ON t.id = ut.task_id
		WHERE ut.user_id = $1
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task testCase.Task
		err := rows.Scan(
			&task.ID,
			&task.Name,
			&task.Description,
			&task.RewardPoints,
		)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
func (r *TaskPostgres) CompleteTask(userId, taskId string) error {
	var userExists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", userId).Scan(&userExists)
	if err != nil {
		return err
	}
	if !userExists {
		return ErrUserNotFound
	}

	var taskExists bool
	err = r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM tasks WHERE id = $1)", taskId).Scan(&taskExists)
	if err != nil {
		return err
	}
	if !taskExists {
		return ErrTaskNotFound
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO user_tasks (user_id, task_id, completed_at) 
		VALUES ($1, $2, NOW())
	`, userId, taskId)
	if err != nil {
		tx.Rollback()
		return err
	}

	var rewardPoints int
	err = tx.QueryRow("SELECT reward_points FROM tasks WHERE id = $1", taskId).Scan(&rewardPoints)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`
		UPDATE users SET balance = balance + $1 WHERE id = $2
	`, rewardPoints, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
