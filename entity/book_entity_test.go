package entity_test

import (
	"book/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithEmptyAuthor(t *testing.T) {
	b := &entity.BookEntity{}
	assert.Panics(t, func() { b.SetAuthor("") }, "The code did panic.")

}
func TestWithEmptyTitle(t *testing.T) {
	b := &entity.BookEntity{}
	assert.Panics(t, func() { b.SetTitle("") }, "The code did panic.")
}

func TestWithEmptyYear(t *testing.T) {
	b := &entity.BookEntity{}
	assert.Panics(t, func() { b.SetYear(0) }, "The code did panic.")
}
