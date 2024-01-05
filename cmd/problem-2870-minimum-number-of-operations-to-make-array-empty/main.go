package main

func main() {
	result := minOperations([]int{
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	})

	println(result)
}

func minOperations(nums []int) int {
	var result int
	var mapping = make(map[int]int)

	const firstOp = 3
	const secondOp = 2

	// build a map of integers
	for _, n := range nums {
		mapping[n]++
	}

	for _, amount := range mapping {
		// edge case: with only 1 unique number
		if amount == 1 {
			return -1
		}

		// simple case: with firstOp
		if (amount % firstOp) == 0 {
			result += amount / firstOp
			continue
		}

		// complex case: with firstOp and 2 * secondOp
		if (amount % firstOp) == 1 {
			result += ((amount / firstOp) - 1) + 2
			continue
		}

		// complex case: with firstOp and a secondOp as the remainder
		if (amount % firstOp) == secondOp {
			result += (amount / firstOp) + 1
			continue
		}

		// simple case: with secondOp
		if (amount % secondOp) == 0 {
			result += amount / secondOp
		}
	}

	return result
}
