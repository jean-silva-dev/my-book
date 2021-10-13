package cli

import (
	"book/usecase/factorymethod"
	"fmt"
)

func Init() {
	var answer int

	fmt.Println("Welcome!")
	fmt.Println("")
	fmt.Println("1 - insert a book;")
	fmt.Println("2 - read a book;")
	fmt.Println("3 - update a book;")
	fmt.Println("4 - remove a book.")

	fmt.Scanln(&answer)

	switch answer {
	case 1:
		insert()
	case 2:
		read()
	case 3:
		update()
	case 4:
		delete()
	}
}

func insert() {
	var title string
	var actor string
	var year int
	fmt.Print("Title name:")
	fmt.Scanln(&title)
	fmt.Print("Actor name:")
	fmt.Scanln(&actor)
	fmt.Print("Publication year:")
	fmt.Scanln(&year)
	buc := factorymethod.CreateBookUseCase()
	buc.CreateBook(title, actor, year)
}

func read() {
	buc := factorymethod.CreateBookUseCase()
	bes := buc.GetBooks()

	fmt.Println("Returned values on database:")
	fmt.Println()
	for _, value := range bes {
		fmt.Printf("Title name: %s. \n", value.Title)
		fmt.Printf("Actor name: %s. \n", value.Author)
		fmt.Printf("Publication year: %d. \n", value.Year)
	}
}

func update() {
	fmt.Println("Not yea implemented.")
}

func delete() {
	fmt.Println("Not yea implemented.")
}
