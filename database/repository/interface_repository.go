package repository

import "book/entity"

type InterfaceRepository interface {
	Insert(entity.BookEntity) error
	Get(entity.BookEntity) error
	Update(entity.BookEntity) error
	Remove(entity.BookEntity) error
}
