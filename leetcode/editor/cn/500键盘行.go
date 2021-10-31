//给你一个字符串数组 words ，只返回可以使用在 美式键盘 同一行的字母打印出来的单词。键盘如下图所示。
//
// 美式键盘 中：
//
//
// 第一行由字符 "qwertyuiop" 组成。
// 第二行由字符 "asdfghjkl" 组成。
// 第三行由字符 "zxcvbnm" 组成。
//
//
//
//
//
//
// 示例 1：
//
//
//输入：words = ["Hello","Alaska","Dad","Peace"]
//输出：["Alaska","Dad"]
//
//
// 示例 2：
//
//
//输入：words = ["omk"]
//输出：[]
//
//
// 示例 3：
//
//
//输入：words = ["adsdf","sfd"]
//输出：["adsdf","sfd"]
//
//
//
//
// 提示：
//
//
// 1 <= words.length <= 20
// 1 <= words[i].length <= 100
// words[i] 由英文字母（小写和大写字母）组成
//
// Related Topics 数组 哈希表 字符串 👍 165 👎 0

package main

import "fmt"

func main() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	fmt.Println(findWords(words))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findWords(words []string) (ret []string) {
	ret = make([]string, 0, len(words))
	set1 := New()
	set2 := New()
	set3 := New()
	set1.AddStr("qwertyuiop")
	set2.AddStr("asdfghjkl")
	set3.AddStr("zxcvbnm")
	for _, v := range words {
		if set1.IsExist(v) || set2.IsExist(v) || set3.IsExist(v) {
			ret = append(ret, v)
		}
	}
	return
}
func (s *set) AddStr(str string) {
	for _, v := range str {
		s.Add(byte(v))
	}
}

func (s *set) IsExist(str string) bool {
	for i := 0; i < len(str); i++ {
		if !s.isExist(str[i]) {
			return false
		}
	}
	return true
}

func (s *set) isExist(word byte) bool {
	if s.Has(word) {
		return true
	}
	if word >= 'a' {
		word -= 32
	} else {
		word += 32
	}
	return s.Has(word)
}

// set 集合
type set struct {
	m map[interface{}]struct{} //字典实现
}

func New() *set {
	return &set{
		m: make(map[interface{}]struct{}),
	}
}
func (s *set) Add(item interface{}) {
	s.m[item] = struct{}{}
}

func (s *set) Remove(item interface{}) {
	delete(s.m, item)
}

func (s *set) Has(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

//leetcode submit region end(Prohibit modification and deletion)
