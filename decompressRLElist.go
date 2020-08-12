package decompress_list

func decompressRLElist(nums []int) []int {
	result := []int{}
	length := len(nums) / 2
	for idx := 0; idx < length; idx++ {
		val := nums[2*idx+1]
		extract := make([]int, nums[2*idx])
		for i := range extract {
			extract[i] = val
		}
		result = append(result, extract...)
	}
	return result
}
