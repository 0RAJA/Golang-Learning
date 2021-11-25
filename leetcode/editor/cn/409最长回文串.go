//给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
//
// 在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
//
// 注意:
//假设字符串的长度不会超过 1010。
//
// 示例 1:
//
//
//输入:
//"abccccdd"
//
//输出:
//7
//
//解释:
//我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
//
// Related Topics 贪心 哈希表 字符串 👍 356 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) (ret int) {
	cnt := [128 - 65]int{}
	for i := range s {
		cnt[s[i]-'A']++ //记录
	}
	k := 0
	for i := range cnt {
		if x := cnt[i] / 2; x > 0 {
			ret += x * 2
		}
		k += cnt[i] % 2
	}
	if k > 0 {
		ret++
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
