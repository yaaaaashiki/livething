package repository

import (
	"github.com/jinzhu/gorm"
)

const Zero = 0

type ObjectRepository struct {
	db *gorm.DB
}

func NewObjectRepository(db *gorm.DB) *ObjectRepository {
	return &ObjectRepository{
		db: db,
	}
}
