package myString

// Kmp s1为主串,s2为模板串,如果s2在s1中出现,返回s1中第一个匹配的字符下标,否则返回-1
func Kmp(s1, s2 string) int {
	if len(s1) < len(s2) {
		return -1
	}
	next := getNext(s2)
	for i, j := 1, 1; i <= len(s1); {
		if j == 0 || s1[i-1] == s2[j-1] {
			i++
			j++
		} else {
			j = next[j]
		}
		if j > len(s2) {
			return i - j
		}
	}
	return -1
}

//返回模板串s的next数组
func getNext(s string) []int {
	next := make([]int, len(s)+1)
	next[1] = 0
	for i, j := 1, 0; i < len(s); {
		if j == 0 || s[i-1] == s[j-1] {
			i++
			j++
			if s[i-1] != s[j-1] {
				next[i] = j
			} else {
				next[i] = next[j]
			}
		} else {
			j = next[j]
		}
	}
	return next
}

// MyCmp 返回s1和s2的最长子序列长度
func MyCmp(s1, s2 string) int {
	dp := make([][]int, len(s1)+1) //行是len(s1)+1
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(s2)+1) //列是len(s2)+1
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = func() int {
					if dp[i-1][j] > dp[i][j-1] {
						return dp[i-1][j]
					}
					return dp[i][j-1]
				}()
			}
		}
	}
	return dp[len(s1)][len(s2)]
}
