package models

type User struct {
	ID            uint   `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Username      string `json:"username" gorm:"not null"`
	Email         string `json:"email" gorm:"not null"`
	FirstName     string `json:"firstName" gorm:"not null"`
	LastName      string `json:"lastName" gorm:"not null"`
	EmailVerified bool   `json:"emailVerified" gorm:"not null"`
	Password      string `json:"-" gorm:"not null"`
}

func (User) TableName() string {
	return "furious-iam.furious-user"
}
