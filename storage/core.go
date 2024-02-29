package storage

import "gorm.io/gorm"

type Storagei interface {
	GetItem(id string, obj any) error
}

func NewGormDB(db *gorm.DB) Storagei {
	return &gormDB{db}
}

func NewDynamoDB() Storagei {
	return &dynamoDB{}
}

// Gorm Database
type gormDB struct {
	DB *gorm.DB
}

func (storage *gormDB) GetItem(id string, obj any) error {
	storage.DB.First(obj, id)
	return nil
}

// Dynamo database
type dynamoDB struct {
}

func (db *dynamoDB) GetItem(id string, obj any) error {
	return nil
}
