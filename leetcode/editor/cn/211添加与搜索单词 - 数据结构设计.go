//请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。
//
// 实现词典类 WordDictionary ：
//
//
// WordDictionary() 初始化词典对象
// void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
// bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回 false 。word 中可能包含一些
//'.' ，每个 . 都可以表示任何一个字母。
//
//
//
//
// 示例：
//
//
//输入：
//["WordDictionary","addWord","addWord","addWord","search","search","search",
//"search"]
//[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
//输出：
//[null,null,null,null,false,true,true,true]
//
//解释：
//WordDictionary wordDictionary = new WordDictionary();
//wordDictionary.addWord("bad");
//wordDictionary.addWord("dad");
//wordDictionary.addWord("mad");
//wordDictionary.search("pad"); // return False
//wordDictionary.search("bad"); // return True
//wordDictionary.search(".ad"); // return True
//wordDictionary.search("b.."); // return True
//
//
//
//
// 提示：
//
//
// 1 <= word.length <= 500
// addWord 中的 word 由小写英文字母组成
// search 中的 word 由 '.' 或小写英文字母组成
// 最多调用 50000 次 addWord 和 search
//
// Related Topics 深度优先搜索 设计 字典树 字符串 👍 334 👎 0

package main

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)

type WordDictionary struct {
	suffix map[byte]*WordDictionary
	values []string
}

func Constructor2() WordDictionary {
	return WordDictionary{suffix: make(map[byte]*WordDictionary)}
}

func (t *WordDictionary) AddWord(k string) {
	root := t
	for i := 0; i < len(k); i++ {
		if root.suffix[k[i]] == nil {
			root.suffix[k[i]] = &WordDictionary{suffix: make(map[byte]*WordDictionary)}
		}
		root = root.suffix[k[i]]
	}
	root.values = append(root.values, "1")
}

func (t *WordDictionary) Search(k string) bool {
	root := t
	for i := 0; i < len(k); i++ {
		if k[i] == '.' {
			for c := 'a'; c <= 'z'; c++ {
				if t.Search(k[:i] + string(c) + k[i+1:]) {
					return true
				}
			}
			return false
		} else {
			if root.suffix[k[i]] != nil {
				root = root.suffix[k[i]]
			} else {
				return false
			}
		}
	}
	return len(root.values) > 0
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
//leetcode submit region end(Prohibit modification and deletion)
