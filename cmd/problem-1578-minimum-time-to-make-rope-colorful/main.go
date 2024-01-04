package main

func main() {
	result := minCost("aaabbbabbbb", []int{3, 5, 10, 7, 5, 3, 5, 5, 4, 8, 1})

	println(result)
}

func minCost(colors string, neededTime []int) int {
	result := 0

	var dc int32 // duplicate color
	var dci int  // duplicate color index cost

	for i, c := range colors {
		if dc != c {
			dc = c
			dci = i

			continue
		}

		if neededTime[i] > neededTime[dci] {
			result = result + neededTime[dci]
			dci = i
		} else {
			result = result + neededTime[i]
		}
	}

	return result
}
