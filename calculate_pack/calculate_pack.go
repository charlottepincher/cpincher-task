package calculate_pack

import (
	"sort"
)

func CalculatePacks(ordered int, pack_sizes []int) map[int]int {
	// This solution assumes that all pack sizes are multiples of the smallest pack size

	// Create a map with each size and quantity
	pack_map := make(map[int]int)
	for i := 0; i < len(pack_sizes); i++ {
		pack_map[pack_sizes[i]] = 0
	}

	// Find the minimum number of items that need to be sent
	min_pack_size := pack_sizes[0]
	var min_items int
	if ordered%min_pack_size != 0 {
		min_items = ((ordered / min_pack_size) + 1) * min_pack_size
	} else {
		min_items = ordered
	}

	rev_pack_sizes := pack_sizes
	sort.Sort(sort.Reverse(sort.IntSlice(rev_pack_sizes)))

	for min_items > 0 {
		for _, size := range rev_pack_sizes {
			if size <= min_items {
				min_items -= size
				pack_map[size] += 1
				break
			}
		}
	}

	return pack_map
}
