package models

import (
	"context"

	"gorm.io/gorm"
)

type DBRepo struct {
	db *gorm.DB
}

func NewDBRepo(db *gorm.DB) DBRepo {
	return DBRepo{db}
}

func (r *DBRepo) FindProduct(id uint64, ctx context.Context) (*Product, error) {
	p, err := gorm.G[Product](r.db).First(ctx)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

type ProductRepo interface {
	FindProduct(id uint64, ctx context.Context) (*Product, error)
	UpdateProduct(id uint64, new Product) (bool, error)
	RemoveProduct(id uint64) error
}

type DishRepo interface {
}

type DataRepo interface {
	ProductRepo
	DishRepo
}
