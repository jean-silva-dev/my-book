package repository

import (
	"book/database/model"
	"book/entity"
)

type InterfaceRepository interface {
	Insert(entity.BookEntity) error
	GetAll() ([]model.BookModel, error)
	Update(entity.BookEntity) error
	Remove(entity.BookEntity) error
}
