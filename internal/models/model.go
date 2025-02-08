package model

import "time"

type Task struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description *string   `gorm:"type:text" json:"description,omitempty"` //опционально
	Status      string    `gorm:"type:varchar(20);not null;check:status IN ('pending', 'in_progress', 'done')" json:"status"`
	Priority    string    `gorm:"type:varchar(10);not null;check:priority IN ('low', 'medium', 'high')" json:"priority"`
	DueDate     time.Time `gorm:"not null" json:"due_date"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar(255);not null" json:"username"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
