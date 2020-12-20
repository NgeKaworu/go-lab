// 316. 去除重复字母
// https://leetcode-cn.com/problems/remove-duplicate-letters/

package removeduplicateletters

func removeDuplicateLetters(s string) string {
	count := [26]int{}
	for _, v := range s {
		count[v-'a']++
	}

	stack := make([]rune, 0)
	inStack := [26]bool{}

	for k, v := range s {
		// 如果当前值不在栈中
		if inStack[v-'a'] == false && k != 0 {
			// 前一个值大于当前值, 且还在后续出现
			for pre := stack[len(stack)-1]; len(stack) > 0 && pre > v && count[pre-'a'] > 0; pre = stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
				inStack[pre-'a'] = false
			}

		}
		// 压栈
		stack = append(stack, v)
		// 标记出现在栈内
		inStack[v-'a'] = true
		// 总数减1
		count[v-'a']--
	}

	return string(stack)
}
