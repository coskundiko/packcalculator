package calculator

import (
	"math"
	"sort"
)

// CalculatePacks finds the best pack combination for a given order.
func CalculatePacks(orderSize int, packSizes []int) map[int]int {
	if orderSize <= 0 {
		return make(map[int]int)
	}

	// We first sort pack sizes to process them in order.
	sort.Ints(packSizes)

	if len(packSizes) == 0 {
		return make(map[int]int)
	}

	// We increase the search limit by + smallest pack size, this is useful if we have a solution more than the order size.
	limit := orderSize + packSizes[0]

	// dp[i] will store the best pack combination for an "i" order size .
	dp := make(map[int]map[int]int)

	// We set up the base case: 0 items, 0 packs
	dp[0] = make(map[int]int)

	// We check every order size up to the limit.
	for i := 1; i < limit; i++ {
		bestPackCount := math.MaxInt32
		var bestCombination map[int]int

		// For each size, we check every available pack.
		for _, pack := range packSizes {

			// We check if we have a solution for the remaining
			subCombination, found := dp[i-pack]
			if i < pack || !found {
				continue
			}

			// If so, we build a new possible solution from the sub-problem.
			combination := make(map[int]int)
			currentPackCount := 0
			for p, c := range subCombination {
				combination[p] = c
				currentPackCount += c
			}
			combination[pack]++
			currentPackCount++

			// If this new solution uses less number of packs, it becomes the new best.
			if currentPackCount < bestPackCount {
				bestPackCount = currentPackCount
				bestCombination = combination
			}
		}

		// We save the best combination found for size i.
		if bestCombination != nil {
			dp[i] = bestCombination
		}

		// Once we find a solution for a size >= order amount, we can return it.
		if i >= orderSize {
			if solution, found := dp[i]; found {
				return solution
			}
		}
	}

	return make(map[int]int)
}
