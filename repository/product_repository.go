package repository

import (
	"github/golang/fiber/multidb/database"
)

type ProductRepository struct {
	db *database.MultiDatabase
}

func NewProduct(db *database.MultiDatabase) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetById() string {
	db, err := r.db.WithTanant("12345")
	{
		if err != nil {
			// ....
		}

		db.Find(1)
	}

	return "data"
}
