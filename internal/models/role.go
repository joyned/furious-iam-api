package models

type Role struct {
	ID          uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name        string `json:"name" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
}

func (Role) TableName() string {
	return "furious-iam.role"
}
