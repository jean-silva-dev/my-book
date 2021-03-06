package usecase

import (
	"book/cli/data"
	"book/database/model"
	"book/database/repository"
	"book/entity"
	"fmt"
)

type BookUseCase struct {
	ir repository.InterfaceRepository
}

func (buc *BookUseCase) SetInterfaceRepository(ir repository.InterfaceRepository) {
	buc.ir = ir
}

func (buc *BookUseCase) CreateBook(bd *data.BookData) {
	be := entity.BookEntity{}
	be.SetTitle(bd.Title)
	be.SetAuthor(bd.Author)
	be.SetYear(bd.Year)

	err := buc.ir.Insert(be)

	if err != nil {
		fmt.Printf("Had a error on create book on database: %s.\n", err)
	}
}

func (buc *BookUseCase) GetBooks() []model.BookModel {
	bms, err := buc.ir.GetAll()

	if err != nil {
		fmt.Printf("Had a error on read book on database: %s.\n", err)
	}

	return bms
}

func (buc *BookUseCase) UpdateBook(bd *data.BookData, index int) {
	be := entity.BookEntity{}
	be.SetTitle(bd.Title)
	be.SetAuthor(bd.Author)
	be.SetYear(bd.Year)
	err := buc.ir.Update(index, be)

	if err != nil {
		fmt.Printf("Had a error on update book on database: %s.\n", err)
	}
}

func (buc *BookUseCase) DeleteBook(index int) {
	err := buc.ir.Remove(index)

	if err != nil {
		fmt.Printf("Had a error on delete book on database: %s.\n", err)
	}
}

func (buc *BookUseCase) FindByNameBook(name string) int {
	index, err := buc.ir.FindByName(name)

	if err != nil {
		fmt.Printf("Had a error on find by name book on database: %s.\n", err)
	}
	return index
}
