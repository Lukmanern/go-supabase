package repository

import "go-supabase/entity"

type TodoRepository interface {
	Create(todo entity.Todo) error

	GetAll() ([]entity.Todo, error)
	GetAllDeleted() ([]entity.Todo, error)
	GetAllNotDeleted() ([]entity.Todo, error)

	Update(id int, todo entity.Todo) error
	UpdateTodoStatus(id int, status string) error

	SoftDeleteById(id int) error
	DeleteById(id int) error

	RestoreTodo(id int) error
}
