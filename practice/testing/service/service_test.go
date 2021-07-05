package service

import (
	"testing"

	"github.com/nitinigi2/go-learning/testing/entity"
)

func TestAdd(t *testing.T) {
	if add(2, 3) != 5 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var tests = []struct {
		x, y     int
		expected int
	}{
		{2, 3, 5},
		{-1, 3, 2},
		{-1, -3, -4},
		{0, 0, 0},
	}

	for _, test := range tests {
		if output := add(test.x, test.y); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, recieved: {}", test.x, test.y, test.expected, output)
		}
	}
}

func TestGetBooks(t *testing.T) {
	books := GetBooks()

	expected := []entity.Book{
		{
			Name:   "book1",
			Author: "author1",
		},
		{
			Name:   "book2",
			Author: "author2",
		},
	}

	if len(books) != len(expected) {
		t.Error("Test Failed")
	}

	for i := 0; i < len(books); i++ {
		if books[i] != expected[i] {
			t.Error("Test Failed")
		}
	}
}
