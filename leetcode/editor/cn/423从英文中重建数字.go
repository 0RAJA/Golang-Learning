//给你一个字符串 s ，其中包含字母顺序打乱的用英文单词表示的若干数字（0-9）。按 升序 返回原始的数字。
//
//
//
// 示例 1：
//
//
//输入：s = "owoztneoer"
//输出："012"
//
//
// 示例 2：
//
//
//输入：s = "fviefuro"
//输出："45"
//
//
//
//
// 提示：
//
//
// 1 <= s.length <= 10⁵
// s[i] 为 ["e","g","f","i","h","o","n","s","r","u","t","w","v","x","z"] 这些字符之一
// s 保证是一个符合题目要求的字符串
//
// Related Topics 哈希表 数学 字符串 👍 130 👎 0

package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(originalDigits("fviefuro"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func originalDigits(s string) (ret string) {
	c := map[rune]int{}
	for _, ch := range s {
		c[ch]++
	}
	cnt := [10]int{}
	//z, w, u, x, g
	cnt[0] = c['z']
	cnt[2] = c['w']
	cnt[4] = c['u']
	cnt[6] = c['x']
	cnt[8] = c['g']
	cnt[3] = c['h'] - cnt[8]
	cnt[5] = c['f'] - cnt[4]
	cnt[7] = c['s'] - cnt[6]
	cnt[1] = c['o'] - cnt[0] - cnt[2] - cnt[4]
	cnt[9] = c['i'] - cnt[5] - cnt[6] - cnt[8]

	ans := []byte{}
	for i, v := range cnt {
		ans = append(ans, bytes.Repeat([]byte{byte('0' + i)}, v)...)
	}
	return string(ans)
}

//leetcode submit region end(Prohibit modification and deletion)
