package usecase_test

import (
	"book/cli/data"
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
	buc := createBookUseCase()
	buc.SetInterfaceRepository(mr)

	/*
		Call function of usecase.
	*/
	bd := createBookData()
	buc.CreateBook(bd)

	mr.AssertExpectations(t)

}

func createBookEntity() entity.BookEntity {
	be := entity.BookEntity{}
	be.SetTitle("The Art of Loving")
	be.SetAuthor("Erich Fromm")
	be.SetYear(1956)
	return be
}

func createBookData() *data.BookData {
	bd := data.BookData{
		Title:  "The Art of Loving",
		Author: "Erich Fromm",
		Year:   1956,
	}
	return &bd
}

func createBookUseCase() *usecase.BookUseCase {
	buc := usecase.BookUseCase{}
	return &buc
}
