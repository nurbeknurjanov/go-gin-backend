package repositories

type PaginationRequest struct {
	PageNumber int `json:"pageNumber"`
	PageSize   int `json:"pageSize"`
}

type PaginationResponse struct {
	PaginationRequest
	PageCount int `json:"pageCount"`
	Total     int `json:"total"`
}

type SortType string

const (
	SORT_ASC  SortType = "asc"
	SORT_DESC SortType = "desc"
)

type Sort struct {
	SortField     string   `json:"sortField"`
	SortDirection SortType `json:"sortDirection"`
}

type List[T any] struct {
	List       *[]T                `json:"list"`
	Pagination *PaginationResponse `json:"pagination"`
}
