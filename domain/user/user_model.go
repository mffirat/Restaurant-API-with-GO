package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	TenantID	  uint  `json:"tenant_id" 		gorm:"not null"`
	Username string `json:"username"        gorm:"unique;not null"`
	Password string `json:"password"        gorm:"not null"`
	Role     string `json:"role"            gorm:"default:'user'"`
}
