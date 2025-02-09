package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Workshop struct {
	ID      uuid.UUID `gorm:"primary_key;type:uuid" json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
}

// Crear workshop
func CreateWorkshop(db *gorm.DB, workshop *Workshop) *Workshop {
	db.Create(workshop)
	return workshop
}

// Obtener todos los workshops
func GetAllWorkshops(db *gorm.DB) []Workshop {
	var workshops []Workshop
	db.Find(&workshops)
	return workshops
}
