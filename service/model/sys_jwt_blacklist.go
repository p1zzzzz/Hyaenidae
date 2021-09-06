package model

import "Hyaenidae/global"

type JwtBlacklist struct {
	global.Hyaenidae_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
