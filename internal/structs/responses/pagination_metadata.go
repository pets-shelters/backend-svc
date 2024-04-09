package responses

type PaginationMetadata struct {
	CurrentPage  int64 `json:"current_page"`
	TotalRecords int64 `json:"total_records"`
}
