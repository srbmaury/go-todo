package todoModels

type Todo struct {
	ID          uint   `gorm:"primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
