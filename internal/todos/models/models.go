package todoModels

import (
	uuid "github.com/satori/go.uuid"
)

type Todo struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
}
