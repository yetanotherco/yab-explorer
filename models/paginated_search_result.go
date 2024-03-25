package models

import "math"

type PaginatedSearchResult struct {
	Page         int         `json:"page"`
	PageSize     int         `json:"pageSize"`
	Results      interface{} `json:"data"`
	ResultsCount int         `json:"resultsCount"`
	PageCount    int         `json:"pageCount"`
}

func NewPaginatedSearchResult(page, pageSize int, results interface{}, resultsCount int) *PaginatedSearchResult {
	pageCount := int(math.Ceil(float64(resultsCount) / float64(pageSize)))

	return &PaginatedSearchResult{
		Page:         page,
		PageSize:     pageSize,
		Results:      results,
		ResultsCount: resultsCount,
		PageCount:    pageCount,
	}
}
