func productExceptSelf(nums []int) []int {
	temp := len(nums)
	result := make([]int, len(nums))

	result[0] = 1
	for i := 1; i < temp; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	right := 1
	for i := temp - 1; i >= 0; i-- {
		result[i] *= right
		right *= nums[i]
	}

	return result
}