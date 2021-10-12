package entity

type BookEntity struct {
	title  string
	author string
	year   int
}

func (be *BookEntity) setTitle(title string) {
	be.title = title
}

func (be *BookEntity) setAuthor(author string) {
	be.author = author
}

func (be *BookEntity) setYear(year int) {
	be.year = year
}
