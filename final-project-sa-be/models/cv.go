package models

import "time"

type CV struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Title     string    `gorm:"size:255;not null" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	Template  string    `gorm:"size:50" json:"template"`
	// Tambahkan relasi many-to-many ke Skill
	Skills    []Skill   `gorm:"many2many:cv_skills;" json:"skills,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
