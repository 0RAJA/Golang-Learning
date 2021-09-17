//给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words，找出所有同时在二维网格和字典中出现的单词。
//
// 单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使
//用。
//
//
//
// 示例 1：
//
//
//输入：board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f",
//"l","v"]], words = ["oath","pea","eat","rain"]
//输出：["eat","oath"]
//
//
// 示例 2：
//
//
//输入：board = [["a","b"],["c","d"]], words = ["abcb"]
//输出：[]
//
//
//
//
// 提示：
//
//
// m == board.length
// n == board[i].length
// 1 <= m, n <= 12
// board[i][j] 是一个小写英文字母
// 1 <= words.length <= 3 * 10⁴
// 1 <= words[i].length <= 10
// words[i] 由小写英文字母组成
// words 中的所有字符串互不相同
//
// Related Topics 字典树 数组 字符串 回溯 矩阵 👍 518 👎 0
package main

import "fmt"

func main() {
	board := [][]byte{{'a', 'b'}, {'c', 'd'}}
	words := []string{"abdc"}
	fmt.Println(findWords(board, words))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findWords(board [][]byte, words []string) (ret []string) {
	next := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	trie := Constructor()
	for i := 0; i < len(words); i++ {
		trie.Insert(words[i], words[i])
	}
	var dfs func(node *Trie, x, y int)
	mapString := make(map[string]bool)
	dfs = func(node *Trie, x, y int) {
		ch := board[x][y]
		node = node.suffix[ch]
		if node == nil {
			return
		}
		if len(node.results) > 0 {
			mapString[node.results[0]] = true
		}
		board[x][y] = '#'
		for i := 0; i < len(next); i++ {
			nx, ny := x+next[i][0], y+next[i][1]
			if nx >= 0 && ny >= 0 && nx < len(board) && ny < len(board[nx]) && board[nx][ny] != '#' {
				dfs(node, nx, ny)
			}
		}
		board[x][y] = ch
	}
	for i := range board {
		for j := range board[i] {
			dfs(trie, i, j)
		}
	}
	ret = make([]string, 0, len(mapString))
	for k := range mapString {
		ret = append(ret, k)
	}
	return
}

type Trie struct {
	suffix  map[byte]*Trie
	results []string
}

func Constructor() *Trie {
	return &Trie{suffix: map[byte]*Trie{}}
}

func (t *Trie) Insert(k, v string) {
	root := t
	for i := 0; i < len(k); i++ {
		if root.suffix[k[i]] == nil {
			root.suffix[k[i]] = &Trie{suffix: make(map[byte]*Trie)}
		}
		root = root.suffix[k[i]]
	}
	root.results = append(root.results, v)
}

func (t *Trie) Search(k string) []string {
	root := t
	for i := 0; i < len(k); i++ {
		if root.suffix[k[i]] != nil {
			root = root.suffix[k[i]]
		} else {
			return []string{}
		}
	}
	return root.results
}

//leetcode submit region end(Prohibit modification and deletion)
