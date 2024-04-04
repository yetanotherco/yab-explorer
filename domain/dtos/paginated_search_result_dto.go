package dtos

import "math"

type PaginatedSearchResultDto struct {
	Page         int         `json:"page"`
	PageSize     int         `json:"pageSize"`
	Results      interface{} `json:"data"`
	ResultsCount int         `json:"resultsCount"`
	PageCount    int         `json:"pageCount"`
}

func NewPaginatedSearchResultDto(page, pageSize int, results interface{}, resultsCount int) *PaginatedSearchResultDto {
	pageCount := int(math.Ceil(float64(resultsCount) / float64(pageSize)))

	return &PaginatedSearchResultDto{
		Page:         page,
		PageSize:     pageSize,
		Results:      results,
		ResultsCount: resultsCount,
		PageCount:    pageCount,
	}
}
