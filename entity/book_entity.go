package entity

import "errors"

type BookEntity struct {
	title  string
	author string
	year   int
}

func (be *BookEntity) GetTitle() string {
	return be.title
}

func (be *BookEntity) GetAuthor() string {
	return be.author
}

func (be *BookEntity) GetYear() int {
	return be.year
}

func (be *BookEntity) SetTitle(title string) error {

	if title == "" {
		return errors.New("the title is empty")
	}

	be.title = title

	return nil
}

func (be *BookEntity) SetAuthor(author string) error {

	if author == "" {
		return errors.New("the author is empty")
	}
	be.author = author

	return nil
}

func (be *BookEntity) SetYear(year int) error {

	if year == 0 {
		return errors.New("the year is not have value")
	}

	be.year = year

	return nil
}
