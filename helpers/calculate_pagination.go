package helpers

import "github.com/marcoagpegoraro/marco_blog/dto"

func CalculatePagination(numberOfPages int, currentPage int, buttonQuantity int) []dto.PaginationButton {

	var paginationButtons []dto.PaginationButton

	if numberOfPages <= buttonQuantity {
		for i := 1; i <= numberOfPages; i++ {
			paginationButtons = append(paginationButtons, dto.PaginationButton{PageNumber: i, IsCurrentPage: i == currentPage})
		}
	} else {
		firstPageToShow := calculateFirstPage(numberOfPages, currentPage, buttonQuantity)

		for i := firstPageToShow; i < firstPageToShow+buttonQuantity; i++ {
			paginationButtons = append(paginationButtons, dto.PaginationButton{PageNumber: i, IsCurrentPage: i == currentPage})
		}
	}

	return paginationButtons
}

func calculateFirstPage(numberOfPages int, currentPage int, buttonQuantity int) int {

	if currentPage == 1 {
		return currentPage
	}

	halfOfButtonQuantity := buttonQuantity / 2

	if currentPage+halfOfButtonQuantity > numberOfPages {
		return numberOfPages - buttonQuantity + 1
	}

	if currentPage > halfOfButtonQuantity {
		return currentPage - halfOfButtonQuantity
	} else {
		return 1 + (halfOfButtonQuantity - currentPage)
	}
}
