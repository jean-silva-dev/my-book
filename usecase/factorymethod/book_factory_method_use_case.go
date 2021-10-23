package factorymethod

import (
	"book/database/repository/factorymethod"
	"book/usecase"
)

func CreateBookUseCase() *usecase.BookUseCase {
	pr := factorymethod.CreatePostgresRepository()
	buc := &usecase.BookUseCase{}
	buc.SetInterfaceRepository(pr)
	return buc
}
