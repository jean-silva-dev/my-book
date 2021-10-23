package usecase_test

import (
	"book/database/repository"
	"book/entity"
	"book/usecase"
	"errors"
	"testing"
)

func TestWithCreateBookSuccessful(t *testing.T) {
	/*
		Create an instance of mock.
	*/
	mr := new(repository.MockRepository)

	/*
		Setup the expectations.
	*/
	be := createBookEntity()
	mr.On("Insert", be).Return(errors.New("Had a error on create book on database."))
	/*
		Setup the usecase.
	*/
	buc := &usecase.BookUseCase{}
	buc.SetInterfaceRepository(mr)

	/*
		Call function of usecase.
	*/
	buc.CreateBook("A", "B", 1990)

	mr.AssertExpectations(t)

}

func createBookEntity() entity.BookEntity {
	be := entity.BookEntity{}
	be.SetAuthor("A")
	be.SetTitle("B")
	be.SetYear(9999)
	return be
}
