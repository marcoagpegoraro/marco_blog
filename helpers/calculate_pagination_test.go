package helpers_test

import (
	"fmt"
	"testing"

	"github.com/marcoagpegoraro/marco_blog/helpers"
)

func TestCalculatePagination1(t *testing.T) {

	numberOfPages := 8
	currentPage := 6
	buttonQuantity := 5

	paginationButtons := helpers.CalculatePagination(numberOfPages, currentPage, buttonQuantity)

	//Expected 4,5,6,7,8

	fmt.Println(paginationButtons)

	if len(paginationButtons) != 5 {
		t.Errorf("Len of buttons is wrong")
	}

	if paginationButtons[0].PageNumber != 4 {
		t.Errorf("First page number incorrect")
	}
}

func TestCalculatePagination2(t *testing.T) {

	numberOfPages := 8
	currentPage := 2
	buttonQuantity := 5

	paginationButtons := helpers.CalculatePagination(numberOfPages, currentPage, buttonQuantity)

	//Expected 1,2,3,4,5

	fmt.Println(paginationButtons)

	if len(paginationButtons) != 5 {
		t.Errorf("Len of buttons is wrong")
	}

	if paginationButtons[0].PageNumber != 1 {
		t.Errorf("First page number incorrect")
	}
}

func TestCalculatePagination3(t *testing.T) {

	numberOfPages := 8
	currentPage := 7
	buttonQuantity := 5

	paginationButtons := helpers.CalculatePagination(numberOfPages, currentPage, buttonQuantity)

	//Expected 4,5,6,7,8

	fmt.Println(paginationButtons)

	if len(paginationButtons) != 5 {
		t.Errorf("Len of buttons is wrong")
	}

	if paginationButtons[0].PageNumber != 4 {
		t.Errorf("First page number incorrect")
	}
}

func TestCalculatePagination4(t *testing.T) {

	numberOfPages := 17
	currentPage := 7
	buttonQuantity := 6

	paginationButtons := helpers.CalculatePagination(numberOfPages, currentPage, buttonQuantity)

	//Expected 4,5,6,7,8,9

	fmt.Println(paginationButtons)

	if len(paginationButtons) != 6 {
		t.Errorf("Len of buttons is wrong")
	}

	if paginationButtons[0].PageNumber != 4 {
		t.Errorf("First page number incorrect")
	}
}

func TestCalculatePagination5(t *testing.T) {

	numberOfPages := 17
	currentPage := 3
	buttonQuantity := 6

	paginationButtons := helpers.CalculatePagination(numberOfPages, currentPage, buttonQuantity)

	//Expected 1,2,3,4,5,6

	fmt.Println(paginationButtons)

	if len(paginationButtons) != 6 {
		t.Errorf("Len of buttons is wrong")
	}

	if paginationButtons[0].PageNumber != 1 {
		t.Errorf("First page number incorrect")
	}
}

func TestCalculatePagination6(t *testing.T) {

	numberOfPages := 17
	currentPage := 17
	buttonQuantity := 6

	paginationButtons := helpers.CalculatePagination(numberOfPages, currentPage, buttonQuantity)

	//Expected 12,13,14,15,16,17

	fmt.Println(paginationButtons)

	if len(paginationButtons) != 6 {
		t.Errorf("Len of buttons is wrong")
	}

	if paginationButtons[0].PageNumber != 12 {
		t.Errorf("First page number incorrect")
	}
}

func TestCalculatePagination7(t *testing.T) {

	numberOfPages := 3
	currentPage := 1
	buttonQuantity := 5

	paginationButtons := helpers.CalculatePagination(numberOfPages, currentPage, buttonQuantity)

	//Expected 1,2,3

	fmt.Println(paginationButtons)

	if len(paginationButtons) != 3 {
		t.Errorf("Len of buttons is wrong")
	}

	if paginationButtons[0].PageNumber != 1 {
		t.Errorf("First page number incorrect")
	}
}
