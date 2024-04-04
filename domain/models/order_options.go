package models

import "slices"

var (
	SortArray      = []string{"order_id", "origin_network", "from_address", "amount", "status", "created_at", "transferred_at", "completed_at"}
	DirectionArray = []string{"asc", "desc"}
)

func SortArrayContains(sort string) bool {
	return slices.Contains(SortArray, sort)
}

func DirectionArrayContains(direction string) bool {
	return slices.Contains(DirectionArray, direction)
}
