// https://leetcode-cn.com/problems/monotone-increasing-digits/
// 738. 单调递增的数字

package monotoneincreasingdigits

import "math"

func monotoneIncreasingDigits(N int) int {
	arr := split(N)
	l := len(arr)
	for i := l - 1; i >= 1; i-- {
		if arr[i] < arr[i-1] {
			for j := i; j < l; j++ {
				arr[j] = 9
			}
			arr[i-1]--
		}
	}
	return join(arr)
}

func split(n int) (res []int) {
	for n != 0 {
		cur := n % 10
		res = append([]int{cur}, res...)
		n /= 10
	}
	return
}

func join(arr []int) (sol int) {
	l := len(arr) - 1
	for k, v := range arr {
		sol += v * int(math.Pow10(l-k))
	}
	return
}
