package repository

import (
	"book/entity"
	"database/sql"
)

type PostgresRepository struct {
	Db *sql.DB
}

func (pr PostgresRepository) Insert(be entity.BookEntity) error {
	_, err := pr.Db.Exec(`
		INSERT INTO books (title, author, year)
		VALUES ($1, $2, $3) `,
		be.GetTitle(),
		be.GetAuthor(),
		be.GetYear())

	return err
}
func (pr PostgresRepository) Get(entity.BookEntity) error {
	return nil
}
func (pr PostgresRepository) Update(entity.BookEntity) error {
	return nil
}
func (pr PostgresRepository) Remove(entity.BookEntity) error {
	return nil
}
