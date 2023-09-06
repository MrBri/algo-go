package removeduplicates

import "sort"

func removeDuplicates(data []int) []int {
	// set := make(map[int]bool)
	// ans := []int{}
	// for _, n := range nums {
	// 	if _, ok := set[n]; ok {
	// 		continue
	// 	}
	// 	set[n] = true
	// }
	// for k := range set {
	// 	ans = append(ans, k)
	// }
	// return ans
	if len(data) <= 1 {
		return data
	}

	// Sort the slice first
	sort.Ints(data)

	// Use two pointers technique to remove duplicates
	i := 0
	for j := 1; j < len(data); j++ {
		if data[i] != data[j] {
			i++
			data[i] = data[j]
		}
	}

	return data[:i+1]
}
