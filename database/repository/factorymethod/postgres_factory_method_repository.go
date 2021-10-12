package factorymethod

import (
	"book/database/connect"
	"book/database/repository"
)

func CreatePostgresRepository() *repository.PostgresRepository {
	db := connect.GetInstance().GetDb()

	pr := &repository.PostgresRepository{
		Db: db}

	return pr
}
