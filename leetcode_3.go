package leetcode_go

import "strings"

//解决思路
//初始最长字串为空，遍历字符数组，若当前字符不在最长字串中，则将其添加到最长字串尾部，更新长度
//否则，记录当前长度，将当前最长字串该重复字符之前的截掉，其余部分作为新字串。
func lengthOfLongestSubstring(s string) int {
	longestSub := ""
	curLen := 0
	maxLen := 0
	for _, i := range []rune(s) {
		if strings.Contains(longestSub, string(i)) {
			//遇到重复字符时
			j := strings.LastIndex(longestSub, string(i))
			longestSub = longestSub[j+1:] + string(i)
			curLen = len(longestSub)
		} else {
			longestSub += string(i)
			curLen++
		}

		if curLen > maxLen {
			maxLen = curLen
		}
	}
	return maxLen
}
