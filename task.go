package testCase

type Task struct {
	ID           string `json:"id" db:"id"`
	Name         string `json:"name" db:"name" binding:"required"`
	Description  string `json:"description" db:"description"`
	RewardPoints int    `json:"reward_points" db:"reward_points" binding:"required"`
}
