package models

// Todos
type Todos struct {
	ID        int    `gorm:"primaryKey" json:"id" example:"1"`
	Title     string `gorm:"type:varchar(255)" json:"title" example:"Belajar Golang"`
	Completed bool   `json:"completed" example:"true"`
}
var Todolist = []Todos{}
