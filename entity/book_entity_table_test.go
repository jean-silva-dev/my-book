package entity_test

import (
	"book/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

type bookEntityTableTest struct {
	function func(value string)
	message  string
}

func TestBookEntityTable(t *testing.T) {
	bett := []bookEntityTableTest{

		bookEntityTableTest{
			function: createBookEntityWithoutAuthor(),
			message:  "The code did panic.",
		},
		bookEntityTableTest{
			function: createBookEntityWithoutTitle(),
			message:  "The code did panic.",
		},
		bookEntityTableTest{
			function: createBookEntityWithoutYear(),
			message:  "The code did panic.",
		},
	}

	for _, value := range bett {
		assert.Panics(t, func() { value.function("") }, value.message)
	}
}

func createBookEntityWithoutTitle() func(value string) {
	return func(value string) {
		be := entity.BookEntity{}
		be.SetTitle(value)
	}
}

func createBookEntityWithoutAuthor() func(value string) {
	return func(value string) {
		be := entity.BookEntity{}
		be.SetTitle(value)
	}
}

func createBookEntityWithoutYear() func(value string) {
	return func(value string) {
		be := entity.BookEntity{}
		be.SetTitle(value)
	}
}
