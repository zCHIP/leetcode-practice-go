package main

func main() {
	result := maxArea([]int{
		1, 8, 6, 2, 5, 4, 8, 3, 7,
	})

	println(result)
}

func maxArea(height []int) int {
	leftP := 0
	rightP := len(height) - 1

	// initial result
	result := calcWater(height[leftP], height[rightP], len(height[leftP:rightP]))

	for leftP < rightP {
		if height[leftP] >= height[rightP] {
			rightP--
		} else {
			leftP++
		}

		tmp := calcWater(height[leftP], height[rightP], len(height[leftP:rightP]))
		if tmp > result {
			result = tmp
		}
	}

	return result
}

func calcWater(leftHeight, rightHeight, numBars int) int {
	if leftHeight > rightHeight {
		return rightHeight * numBars
	}

	if rightHeight > leftHeight {
		return leftHeight * numBars
	}

	return leftHeight * numBars
}
