package calculate_pack

import (
	"sort"
)

func CalculatePacks(ordered int, pack_sizes []int, gcd int) map[int]int {
	var possible_packs [][]int
	if ordered < pack_sizes[0] {
		var smallest_pack []int
		smallest_pack = append(smallest_pack, pack_sizes[0])
		possible_packs = append(possible_packs, smallest_pack)
	} else {
		remainder := ordered % gcd
		if remainder != 0 {
			ordered += (gcd - remainder)
		}
		possible_packs = PossiblePacks(pack_sizes, ordered)
		for possible_packs == nil {
			ordered += 1
			possible_packs = PossiblePacks(pack_sizes, ordered)
		}
	}
	answer := FindSolution(possible_packs, pack_sizes)
	return answer
}

func FindGCD(pack_sizes []int) int {
	gcd := GCDEuclidean(pack_sizes[0], pack_sizes[1])
	for ii := 2; ii < len(pack_sizes); ii++ {
		gcd = GCDEuclidean(gcd, pack_sizes[ii])
	}
	return gcd
}

func GCDEuclidean(num1, num2 int) int {
	for num1 != num2 {
		if num1 > num2 {
			num1 -= num2
		} else {
			num2 -= num1
		}
	}
	return num1
}

func PossiblePacks(pack_sizes []int, ordered int) [][]int {
	var possible_packs [][]int
	var current_packs []int
	sort.Ints(pack_sizes)
	possible_packs = FindPossibilities(possible_packs, pack_sizes, current_packs, ordered, 0)
	return possible_packs
}

func FindPossibilities(possible_packs [][]int, pack_sizes []int, current_packs []int, ordered int, index int) [][]int {
	if ordered == 0 {
		possible_packs = append(possible_packs, current_packs)
		return possible_packs
	}
	for ii := index; ii < len(pack_sizes); ii++ {
		pack := pack_sizes[ii]
		if ordered-pack >= 0 {
			current_packs = append(current_packs, pack)
			possible_packs = FindPossibilities(possible_packs, pack_sizes, current_packs, (ordered - pack), ii)
			current_packs = current_packs[:len(current_packs)-1]
		}
	}
	return possible_packs
}

func FindSolution(possible_packs [][]int, pack_sizes []int) map[int]int {
	// Find the smallest amount of packs
	smallest_len := len(possible_packs[0])
	smallest_packs := possible_packs[0]
	for _, packs := range possible_packs {
		if len(packs) < smallest_len {
			smallest_len = len(packs)
			smallest_packs = packs
		}
	}
	pack_map := make(map[int]int)
	for i := 0; i < len(pack_sizes); i++ {
		pack_map[pack_sizes[i]] = 0
	}
	for _, pack := range smallest_packs {
		pack_map[pack] += 1
	}
	return pack_map
}
