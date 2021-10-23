package repository

import (
	"book/database/model"
	"book/entity"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mr MockRepository) Insert(be entity.BookEntity) error {
	args := mr.Called(be)
	return args.Error(0)
}

func (mr *MockRepository) GetAll() ([]model.BookModel, error) {
	args := mr.Called()
	return (make([]model.BookModel, 0)), args.Error(1)
}

func (mr *MockRepository) Update(index int, be entity.BookEntity) error {
	args := mr.Called(index, be)
	return args.Error(0)
}

func (mr *MockRepository) Remove(index int) error {
	args := mr.Called(index)
	return args.Error(0)
}

func (mr *MockRepository) FindByName(name string) (int, error) {
	args := mr.Called(name)
	return args.Int(0), args.Error(1)
}
