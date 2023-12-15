package credentials

import "gorm.io/gorm"

type Credentials struct {
	gorm.Model
	ID           uint   `json:"id" gorm:"AUTO_INCREMENT;"primaryKey"`
	Username     string `json:"username" gorm:"not null"`
	Password     string `json:"password" gorm:"not null"`
	ClientSecret string `json:"client_secret" gorm:"not null"`
	ClientID     string `json:"client_id" gorm:"not null"`
}
