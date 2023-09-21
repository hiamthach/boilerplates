package common

type Pagination struct {
	Total       int64 `json:"total"`
	CurrentPage int64 `json:"current_page"`
	PerPage     int64 `json:"per_page"`
}
