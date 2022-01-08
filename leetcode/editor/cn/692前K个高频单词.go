//给一非空的单词列表，返回前 K 个出现次数最多的单词。
//
// 返回的答案应该按单词出现频率由高到低排序。如果不同的单词有相同出现频率，按字母顺序排序。
//
// 示例 1：
//
//
//输入: ["i", "love", "leetcode", "i", "love", "coding"], K = 2
//输出: ["i", "love"]
//解析: "i" 和 "love" 为出现次数最多的两个单词，均为2次。
//    注意，按字母顺序 "i" 在 "love" 之前。
//
//
//
//
// 示例 2：
//
//
//输入: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], K
// = 4
//输出: ["the", "is", "sunny", "day"]
//解析: "the", "is", "sunny" 和 "day" 是出现次数最多的四个单词，
//    出现次数依次为 4, 3, 2 和 1 次。
//
//
//
//
// 注意：
//
//
// 假定 K 总为有效值， 1 ≤ K ≤ 集合元素数。
// 输入的单词均由小写字母组成。
//
//
//
//
// 扩展练习：
//
//
// 尝试以 O(n log K) 时间复杂度和 O(n) 空间复杂度解决。
//
// Related Topics 堆 字典树 哈希表
// 👍 291 👎 0

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	words := []string{"the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"}
	fmt.Println(topKFrequent(words, 4))
}

//leetcode submit region begin(Prohibit modification and deletion)
type myType []string

var map1 map[string]int

func topKFrequent(words []string, k int) []string {
	map1 = make(map[string]int)
	for _, v := range words {
		map1[v]++
	}
	key := make(myType, 0)
	for k := range map1 {
		key = append(key, k)
	}
	sort.Sort(key)
	return key[:k]
}
func (m myType) Len() int {
	return len(m)
}
func (m myType) Less(i, j int) bool {
	if map1[m[i]] != map1[m[j]] {
		return map1[m[i]] > map1[m[j]]
	} else {
		return strings.Compare(m[i], m[j]) < 0
	}
}
func (m myType) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

//leetcode submit region end(Prohibit modification and deletion)
