package models

type Skill struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:100;not null;unique" json:"name"`
	// Hubungan many-to-many dengan CV
	CVs []CV `gorm:"many2many:cv_skills;" json:"cvs,omitempty"`
}
