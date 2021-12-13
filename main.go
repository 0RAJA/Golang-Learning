package main

import "fmt"

func main() {
	//preorder := []int{3, 9, 20, 15, 7}
	//inorder := []int{9, 3, 15, 20, 7}
	//buildTree(preorder, inorder)
	fmt.Println([]int{} == nil)
}

func exist(board [][]byte, word string) bool {
	nextx := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	startPoint := [][2]int{}
	for i := range board {
		for j := range board[i] {
			if word[0] == board[i][j] {
				startPoint = append(startPoint, [2]int{i, j})
			}
		}
	}
	maps := make([][]bool, len(board))
	for i := range maps {
		maps[i] = make([]bool, len(board[i]))
	}
	var dfs func(sx, sy, next int) bool
	dfs = func(x, y, next int) bool {
		if next >= len(word) {
			return true
		}
		maps[x][y] = true
		for i := 0; i < 4; i++ {
			nx, ny := x+nextx[i][0], y+nextx[i][1]
			if nx < 0 || ny < 0 || nx >= len(board) || ny >= len(board[0]) || maps[nx][ny] || board[nx][ny] != word[next] {
				continue
			}
			maps[nx][ny] = true
			if dfs(nx, ny, next+1) == true {
				return true
			}
			maps[nx][ny] = false
		}
		return false
	}
	for i := range startPoint {
		if ok := dfs(startPoint[i][0], startPoint[i][1], 1); ok {
			return true
		}
	}
	return false
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var dfs func(preorder, inorder []int) *TreeNode
	dfs = func(preorder, inorder []int) (ret *TreeNode) {
		if preorder == nil {
			return nil
		}
		k := preorder[0]
		ret = new(TreeNode)
		ret.Val = k
		index := 0
		for inorder[index] != k {
			index++
		}
		if index != 0 {
			ret.Left = dfs(preorder[1:index+1], inorder[:index])
		}
		if index != len(inorder)-1 {
			ret.Right = dfs(preorder[index+1:], inorder[index+1:])
		}
		return
	}
	return dfs(preorder, inorder)
}
