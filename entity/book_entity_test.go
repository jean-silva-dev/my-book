package entity_test

import (
	"book/entity"
	"testing"
)

func TestWithEmptyTitle(t *testing.T) {
	b := &entity.BookEntity{}
	actual := b.SetTitle("")

	if actual == nil {
		t.Fatal("Expected:", nil, "got:", actual)
	}

}

func TestWithEmptyAuthor(t *testing.T) {
	b := &entity.BookEntity{}
	actual := b.SetAuthor("")

	if actual == nil {
		t.Fatal("Expected:", nil, "got:", actual)
	}

}

func TestWithEmptyYear(t *testing.T) {
	b := &entity.BookEntity{}
	actual := b.SetYear(0)

	if actual == nil {
		t.Fatal("Expected:", nil, "got:", actual)
	}

}
