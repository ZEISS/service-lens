package models

// Pagination is a struct that contains the pagination information.
type Pagination struct {
	// Limit is the number of items to return.
	Limit int `json:"limit"`
	// Offset is the number of items to skip.
	Offset int `json:"offset"`
}
