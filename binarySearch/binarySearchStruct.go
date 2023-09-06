package binarysearch

import "math"

func BinarySearchStruct(data []struct {
	name      string
	followers int
}, followers int,
) string {
	l := 0
	r := len(data) - 1
	for l <= r {
		mid := int(math.Round(float64(l+r) / 2))
		if data[mid].followers == followers {
			return data[mid].name
		} else if data[mid].followers < followers {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ""
}
