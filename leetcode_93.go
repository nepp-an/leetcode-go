package leetcode_go

import (
	"fmt"
	"strconv"
)

const SEG_COUNT = 4

var (
	ans      []string
	segments []int
)

func restoreIpAddresses(s string) []string {
	segments = make([]int, SEG_COUNT)
	ans = []string{}
	dfs(s, 0, 0)
	return ans
}

func dfs(s string, segId, segStart int) {
	// 如果找到了 4 段 IP 地址并且遍历完了字符串，那么就是一种答案
	if segId == SEG_COUNT {
		if segStart == len(s) {
			ipAddr := ""
			for i := 0; i < SEG_COUNT; i++ {
				ipAddr += strconv.Itoa(segments[i])
				if i != SEG_COUNT-1 {
					ipAddr += "."
				}
			}
			ans = append(ans, ipAddr)
		}
		return
	}

	// 如果还没有找到 4 段 IP 地址就已经遍历完了字符串，那么提前回溯
	if segStart == len(s) {
		return
	}
	// 由于不能有前导零，如果当前数字为 0，那么这一段 IP 地址只能为 0
	if s[segStart] == '0' {
		segments[segId] = 0
		dfs(s, segId+1, segStart+1)
	}
	// 一般情况，枚举每一种可能性并递归
	addr := 0
	for segEnd := segStart; segEnd < len(s); segEnd++ {
		addr = addr*10 + int(s[segEnd]-'0')
		if addr > 0 && addr <= 0xFF {
			segments[segId] = addr
			dfs(s, segId+1, segEnd+1)
		} else {
			break
		}
	}
}

func main() {
	var test = "25525511135"
	fmt.Println(restoreIpAddresses(test))
}

/*
func restoreIpAddresses(s string) []string {
    if length:=len(s); length> 12 ||length < 4  {
        return nil
    }
	ans := make([]string, 0)
	for i := 1; i < len(s) - 2; i++ {
	two:
		for j := i + 1; j < len(s) - 1; j++ {
			for k := j + 1; k < len(s); k++ {
				var val1, val2, val3, val4 string
				var ok bool
				if val1, ok = check(s[:i]); !ok {
					return ans
				}
				if val2, ok = check(s[i:j]); !ok {
					break two
				}
				if val3, ok = check(s[j:k]); !ok {
					break
				}
				if val4, ok = check(s[k:]); ok {
					ans = append(ans, val1+"."+val2+"."+val3+"."+val4)
				}
			}
		}
	}
	return ans
}

func check(s string) (string, bool) {
    if len(s) == 0 || (len(s) > 1 &&  s[0] == '0') {
       return "", false
    }
	var val int
	if val, _ = strconv.Atoi(s); val > 255{
		return "", false
	}
	return s, true
}
*/
