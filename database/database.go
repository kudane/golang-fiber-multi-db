package database

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type MultiDatabase struct {
	connectors map[string]*gorm.DB
}

func New() *MultiDatabase {
	multidb := &MultiDatabase{
		connectors: map[string]*gorm.DB{
			"tanant1": &gorm.DB{},
			"tanant2": &gorm.DB{},
		},
	}

	return multidb
}

func (d *MultiDatabase) WithTanant(tanantName string) (*gorm.DB, error) {
	// check tanant, not found throw error
	return d.connectors["tanantName"], nil
	// return nil, errors.New("no tanant")
}

func (d *MultiDatabase) WithTanant(c *fiber.Ctx) any {
	ctx := context.Background()
	tanantName := ctx.Value("tanantId").(string)
	return d.connectors[tanantName]
}

func (d *MultiDatabase) Tanant2(tanantName string) any {
	return d.connectors[tanantName]
}
