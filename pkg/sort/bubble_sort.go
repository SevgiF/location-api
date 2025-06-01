package sort

import "github.com/SevgiF/location-api/internal/core/location/domain"

// SortingDistanceWithBubbleSort for sorting locations by distance
func SortingDistanceWithBubbleSort(arr []domain.LocationRouting, size int) []domain.LocationRouting {
	if size == 1 {
		return arr
	}

	var i = 0
	for i < size-1 {
		if arr[i].Distance > arr[i+1].Distance {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
		i++
	}

	SortingDistanceWithBubbleSort(arr, size-1)

	return arr
}
