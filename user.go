package testCase

type User struct {
	ID         string  `json:"id" db:"id"`
	UserName   string  `json:"username" db:"username" binding:"required"`
	Password   string  `json:"password" db:"password_hash" binding:"required"`
	Balance    int     `json:"balance" db:"balance"`
	ReferrerID *string `json:"referrer_id,omitempty" db:"referrer_id"`
}
