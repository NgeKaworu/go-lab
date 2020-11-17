package lengthoflis

func lengthOfLIS(nums []int) int {
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		tmp := 1
		for j := i + 1; j < n; j++ {
			if nums[j] > nums[i] && nums[j] > nums[j-1] {
				tmp++
			}
		}
		res = max(res, tmp)
	}
	return res
}

func max(nums ...int) int {
	res := nums[0]
	for _, v := range nums {
		if v > res {
			res = v
		}
	}
	return res
}
