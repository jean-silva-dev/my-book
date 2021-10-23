package entity

type BookEntity struct {
	author string
	title  string
	year   int
}

func (be *BookEntity) GetAuthor() string {
	return be.author
}

func (be *BookEntity) GetTitle() string {
	return be.title
}

func (be *BookEntity) GetYear() int {
	return be.year
}

func (be *BookEntity) SetAuthor(author string) {

	if author == "" {
		panic("Author cannot be empty.")
	}
	be.author = author
}

func (be *BookEntity) SetTitle(title string) {

	if title == "" {
		panic("Title cannot be empty.")
	}
	be.title = title
}

func (be *BookEntity) SetYear(year int) {

	if year == 0 {
		panic("Year cannot be 0.")
	}
	be.year = year
}
