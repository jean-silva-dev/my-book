package cli

import (
	"book/cli/data"
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

		// answer = 1
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
	fmt.Println("")
	fmt.Println("Insert data from new book.")
	fmt.Println("")
	be := createBookEntity()
	buc := factorymethod.CreateBookUseCase()
	buc.CreateBook(be)
	log.Println("Created with successful.")
}

func read() {
	books := getAllBooks()

	if len(books) == 0 {
		log.Println("None data was found on database.")
	} else {
		fmt.Println("")
		log.Println("Returned values on database:")
		fmt.Println("")
		fmt.Println("----------")

		for _, value := range books {
			fmt.Printf("Actor name: %s. \n", value.Author)
			fmt.Printf("Title name: %s. \n", value.Title)
			fmt.Printf("Publication year: %d. \n", value.Year)
			fmt.Println("----------")
		}
	}
}

func update() {
	books := getAllBooks()

	if len(books) == 0 {
		return
	}

	if len(books) == 1 {
		fmt.Println("")
		fmt.Println("Do insert new dates, please.")
		fmt.Println("")
		be := createBookEntity()

		element := books[0]
		buc := factorymethod.CreateBookUseCase()
		indexForChange := buc.FindByNameBook(element.Title)
		buc.UpdateBook(be, indexForChange)
		log.Println("Updated with successful.")
	} else {
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

			fmt.Println("")
			fmt.Println("Do insert new dates, please.")
			fmt.Println("")
			be := createBookEntity()

			buc.UpdateBook(be, indexForChange)
			log.Println("Updated with successful.")
			break
		}
	}
}

func delete() {
	books := getAllBooks()

	if len(books) == 0 {
		log.Println("None data was found on database.")
	} else {

		for {
			fmt.Println("")
			fmt.Println("What is the item that you wish to delete?")
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

			buc.DeleteBook(indexForChange)
			log.Println("Removed with successful.")
			break
		}
	}
}

func createBookEntity() *data.BookData {
	bd := &data.BookData{}

	fmt.Print("Title name: ")
	fmt.Scanln(&bd.Title)
	fmt.Print("Actor name: ")
	fmt.Scanln(&bd.Author)
	fmt.Print("Publication year: ")
	fmt.Scanln(&bd.Year)
	return bd
}

func getAllBooks() []model.BookModel {
	buc := factorymethod.CreateBookUseCase()
	bes := buc.GetBooks()
	return bes
}
