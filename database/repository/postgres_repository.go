package repository

import (
	"book/database/model"
	"book/entity"
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

func (pr *PostgresRepository) SetDb(db *sql.DB) {
	pr.db = db
}

func (pr *PostgresRepository) Insert(be entity.BookEntity) error {
	_, err := pr.db.Exec(`
		INSERT INTO books (title, author, year)
		VALUES ($1, $2, $3) `,
		be.GetTitle(),
		be.GetAuthor(),
		be.GetYear())

	return err
}
func (pr *PostgresRepository) GetAll() ([]model.BookModel, error) {
	bms := make([]model.BookModel, 0)

	rows, err := pr.db.Query(`SELECT title, author, year FROM books`)

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
func (pr *PostgresRepository) Update(index int, be entity.BookEntity) error {
	_, err := pr.db.Exec(`
		UPDATE books
		SET title = $1, author = $2, year = $3
		WHERE id = $4
		`,
		be.GetTitle(),
		be.GetAuthor(),
		be.GetYear(),
		index)

	return err
}

func (pr *PostgresRepository) Remove(index int) error {
	_, err := pr.db.Exec(`
		DELETE FROM books
		WHERE id = $1
		`,
		index)

	return err
}
func (pr *PostgresRepository) FindByName(name string) (int, error) {
	var index int

	row := pr.db.QueryRow(`
		SELECT id
		FROM books
		WHERE title = $1
		`,
		name)

	err := row.Scan(
		&index,
	)

	return index, err
}
