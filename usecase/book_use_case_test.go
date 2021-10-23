package usecase_test

import (
	"book/database/repository"
	"book/usecase"
	"errors"
	"testing"
)

func TestWithCreateBookSuccessful(t *testing.T) {
	mr := new(repository.MockRepository)
	mr.On("Insert").Return(errors.New("a"))
	buc := &usecase.BookUseCase{
		Ir: mr,
	}

	buc.CreateBook("a", "b", 3)
	mr.AssertExpectations(t)
}
