package usecase

import "book/database/repository"

type BookUseCase struct {
	ir repository.InterfaceRepository
}

func (buc BookUseCase) createBook() {
	buc.ir.Insert()
}
