package cli

import (
	"book/database/model"
	"book/usecase/factorymethod"
	"fmt"
	"log"
)

func Init() {

	for {
		var answer int
		fmt.Println("")
		fmt.Println("Welcome!")
		fmt.Println("")
		fmt.Println("1 - insert a book;")
		fmt.Println("2 - read a book;")
		fmt.Println("3 - update a book;")
		fmt.Println("4 - remove a book;")
		fmt.Println("0 - exit programm.")

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
		case 0:
			return
		default:
			log.Println("Please, select a right value.")
		}
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
	books := getAllBooks()

	fmt.Println("")
	fmt.Println("Returned values on database:")
	fmt.Println("")
	fmt.Println("----------")
	for _, value := range books {
		fmt.Printf("Actor name: %s. \n", value.Author)
		fmt.Printf("Title name: %s. \n", value.Title)
		fmt.Printf("Publication year: %d. \n", value.Year)
		fmt.Println("----------")
	}
}

func update() {
	books := getAllBooks()

	for {
		fmt.Println("")
		fmt.Println("What is the item that you wish to change?")
		fmt.Println("")
		fmt.Println("----------")
		for index, value := range books {
			fmt.Printf("[%d]\n", index+1)
			fmt.Printf("Actor name: %s. \n", value.Author)
			fmt.Printf("Title name: %s. \n", value.Title)
			fmt.Printf("Publication year: %d. \n", value.Year)
			fmt.Println("----------")
		}
		var answer int
		fmt.Scanln(&answer)

		if answer <= 0 {
			log.Println("Please, select a right value.")
			continue
		}

		if answer > len(books) {
			log.Println("Please, select a right value.")
			continue
		}

		index := answer - 1
		element := books[index]
		buc := factorymethod.CreateBookUseCase()
		indexForChange := buc.FindByNameBook(element.Title)

		var title string
		var actor string
		var year int
		fmt.Println("Do insert new dates, please.")
		fmt.Print("Title name:")
		fmt.Scanln(&title)
		fmt.Print("Actor name:")
		fmt.Scanln(&actor)
		fmt.Print("Publication year:")
		fmt.Scanln(&year)

		buc.UpdateBook(title, actor, year, indexForChange)
		break
	}
}

func delete() {
	fmt.Println("Not yea implemented.")
}

func getAllBooks() []model.BookModel {
	buc := factorymethod.CreateBookUseCase()
	bes := buc.GetBooks()
	return bes
}
