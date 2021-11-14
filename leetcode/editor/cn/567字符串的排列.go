//给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
//
// 换句话说，s1 的排列之一是 s2 的 子串 。
//
//
//
// 示例 1：
//
//
//输入：s1 = "ab" s2 = "eidbaooo"
//输出：true
//解释：s2 包含 s1 的排列之一 ("ba").
//
//
// 示例 2：
//
//
//输入：s1= "ab" s2 = "eidboaoo"
//输出：false
//
//
//
//
// 提示：
//
//
// 1 <= s1.length, s2.length <= 10⁴
// s1 和 s2 仅包含小写字母
//
// Related Topics 哈希表 双指针 字符串 滑动窗口 👍 489 👎 0

package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	cntMap1 := [26]int{} //记录s1中字符出现次数
	eleCnt := 0          //s1中在s2中未出现的字符的个数
	for i := range s1 {
		if cntMap1[s1[i]-'a'] == 0 {
			eleCnt++
		}
		cntMap1[s1[i]-'a']++
	}
	for l, r := 0, 0; r < len(s2); r++ {
		cntMap1[s2[r]-'a']-- //尝试通过s2中添加的新元素抵消s1中的出现的元素,如果能完全抵消,则减小elecnt个数
		if cntMap1[s2[r]-'a'] == 0 {
			eleCnt--
		}
		if r-l+1 > len(s1) { //尝试通过淘汰的元素恢复s1中字符的出现次数,如果能恢复出来,则增加elecnt个数
			if cntMap1[s2[l]-'a'] == 0 {
				eleCnt++
			}
			cntMap1[s2[l]-'a']++
			l++
		}
		if eleCnt == 0 {
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
