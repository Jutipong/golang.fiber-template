package model

import "gorm.io/gorm"

type ConnectDB struct {
	Test *gorm.DB
	Demo *gorm.DB
}
