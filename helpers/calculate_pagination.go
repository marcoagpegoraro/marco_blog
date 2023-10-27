package helpers

import "github.com/marcoagpegoraro/marco_blog/dto"

func CalculatePagination(numberOfPages int, currentPage int, desiredButtonQuantity int) []dto.PaginationButton {

	var paginationButtons []dto.PaginationButton

	if numberOfPages <= desiredButtonQuantity {
		for i := 1; i <= numberOfPages; i++ {
			paginationButtons = append(paginationButtons, dto.PaginationButton{PageNumber: i, IsCurrentPage: i == currentPage})
		}
	} else {
		firstPageToShow := calculateFirstPage(numberOfPages, currentPage, desiredButtonQuantity)

		for i := firstPageToShow; i < firstPageToShow+desiredButtonQuantity; i++ {
			paginationButtons = append(paginationButtons, dto.PaginationButton{PageNumber: i, IsCurrentPage: i == currentPage})
		}
	}

	return paginationButtons
}

func calculateFirstPage(numberOfPages int, currentPage int, desiredButtonQuantity int) int {

	if currentPage == 1 {
		return currentPage
	}

	halfOfButtonQuantity := desiredButtonQuantity / 2

	if currentPage+halfOfButtonQuantity > numberOfPages {
		return numberOfPages - desiredButtonQuantity + 1
	}

	if currentPage > halfOfButtonQuantity {
		return currentPage - halfOfButtonQuantity
	} else {
		return 1 + (halfOfButtonQuantity - currentPage)
	}
}
