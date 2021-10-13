package usecase

import (
	"book/database/model"
	"book/database/repository"
	"book/entity"
	"fmt"
)

type BookUseCase struct {
	Ir repository.InterfaceRepository
}

func (buc BookUseCase) CreateBook(title, actor string, year int) {
	be := entity.BookEntity{}
	be.SetTitle(title)
	be.SetAuthor(actor)
	be.SetYear(year)
	err := buc.Ir.Insert(be)

	if err != nil {
		fmt.Printf("Had a error on create book on database: %s.\n", err)
	}
}

func (buc BookUseCase) GetBooks() []model.BookModel {
	bms, err := buc.Ir.GetAll()

	if err != nil {
		fmt.Printf("Had a error on read book on database: %s.\n", err)
	}

	return bms
}
