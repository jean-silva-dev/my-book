package entity_test

import (
	"book/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueWithEmptyAuthor(t *testing.T) {
	b := &entity.BookEntity{}
	assert.Panics(t, func() { b.SetAuthor("") }, "The code did panic.")

}
func TestUniqueWithEmptyTitle(t *testing.T) {
	b := &entity.BookEntity{}
	assert.Panics(t, func() { b.SetTitle("") }, "The code did panic.")
}

func TestUniqueWithEmptyYear(t *testing.T) {
	b := &entity.BookEntity{}
	assert.Panics(t, func() { b.SetYear(0) }, "The code did panic.")
}
