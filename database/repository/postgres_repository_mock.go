package repository

import (
	"book/database/model"
	"book/entity"
	"fmt"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mr MockRepository) Insert(entity.BookEntity) error {
	args := mr.Called()
	result := args.Get(0)
	fmt.Println(result)
	return args.Error(0)
}

func (mr MockRepository) GetAll() ([]model.BookModel, error) {
	args := mr.Called()
	// result := args.Get(0)
	return (make([]model.BookModel, 0)), args.Error(1)
}
func (mr MockRepository) Update(entity.BookEntity) error {
	args := mr.Called()
	// result := args.Get(0)
	return args.Error(1)
}
func (mr MockRepository) Remove(entity.BookEntity) error {
	args := mr.Called()
	// result := args.Get(0)
	return args.Error(1)
}
