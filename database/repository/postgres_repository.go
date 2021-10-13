package repository

import (
	"book/database/model"
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
func (pr PostgresRepository) GetAll() ([]model.BookModel, error) {
	bms := make([]model.BookModel, 0)

	rows, err := pr.Db.Query(`SELECT title, author, year FROM books`)

	if err != nil {
		return bms, err
	}
	defer rows.Close()

	for rows.Next() {
		var bm model.BookModel

		err := rows.Scan(
			&bm.Title,
			&bm.Author,
			&bm.Year,
		)

		if err != nil {
			return bms, err
		}
		bms = append(bms, bm)
	}

	return bms, err
}
func (pr PostgresRepository) Update(entity.BookEntity) error {
	return nil
}
func (pr PostgresRepository) Remove(entity.BookEntity) error {
	return nil
}
