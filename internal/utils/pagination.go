package utils

func CalculatePagination(totalItems, page, pageSize int) (int, int) {
	totalPages := (totalItems + pageSize - 1) / pageSize
	if page > totalPages {
		page = totalPages
	}

	return totalItems, totalPages
}
