package models

type Class struct {
	ID          uint   `gorm:"primaryKey"`
	ClassID     string `gorm:"uniqueIndex:idx_class_id;size:10"`
	Name        string
	CreatedById uint
	User        User `gorm:"foreignKey:CreatedById"`
}
