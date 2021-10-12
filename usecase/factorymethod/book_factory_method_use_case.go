package factorymethod

import (
	"book/database/repository/factorymethod"
	"book/usecase"
)

func CreateBookUseCase() *usecase.BookUseCase {
	pr := factorymethod.CreatePostgresRepository()
	buc := &usecase.BookUseCase{
		Ir: pr,
	}
	return buc
}
