//在一个 N x N 的坐标方格 grid 中，每一个方格的值 grid[i][j] 表示在位置 (i,j) 的平台高度。
//
// 现在开始下雨了。当时间为 t 时，此时雨水导致水池中任意位置的水位为 t 。你可以从一个平台游向四周相邻的任意一个平台，但是前提是此时水位必须同时淹没这两
//个平台。假定你可以瞬间移动无限距离，也就是默认在方格内部游动是不耗时的。当然，在你游泳的时候你必须待在坐标方格里面。
//
// 你从坐标方格的左上平台 (0，0) 出发。最少耗时多久你才能到达坐标方格的右下平台 (N-1, N-1)？
//
//
//
// 示例 1:
//
//
//输入: [[0,2],[1,3]]
//输出: 3
//解释:
//时间为0时，你位于坐标方格的位置为 (0, 0)。
//此时你不能游向任意方向，因为四个相邻方向平台的高度都大于当前时间为 0 时的水位。
//
//等时间到达 3 时，你才可以游向平台 (1, 1). 因为此时的水位是 3，坐标方格中的平台没有比水位 3 更高的，所以你可以游向坐标方格中的任意位置
//
//
// 示例2:
//
//
//输入: [[0,1,2,3,4],[24,23,22,21,5],[12,13,14,15,16],[11,17,18,19,20],[10,9,8,7,6
//]]
//输出: 16
//解释:
// 0  1  2  3  4
//24 23 22 21  5
//12 13 14 15 16
//11 17 18 19 20
//10  9  8  7  6
//
//最终的路线用加粗进行了标记。
//我们必须等到时间为 16，此时才能保证平台 (0, 0) 和 (4, 4) 是连通的
//
//
//
//
// 提示:
//
//
// 2 <= N <= 50.
// grid[i][j] 是 [0, ..., N*N - 1] 的排列。
//
// Related Topics 深度优先搜索 广度优先搜索 并查集 数组 二分查找 矩阵 堆（优先队列） 👍 213 👎 0

package main

import (
	"fmt"
)

func main() {
	grid := [][]int{{3, 2}, {1, 0}}
	fmt.Println(swimInWater(grid))
}

//leetcode submit region begin(Prohibit modification and deletion)
type Set struct {
	parent, size []int
}

func NewSet(n int) *Set {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range size {
		parent[i] = i
		size[i] = 1
	}
	return &Set{parent: parent, size: size}
}

func (s *Set) find(x int) int {
	if s.parent[x] != x {
		s.parent[x] = s.find(s.parent[x])
	}
	return s.parent[x]
}

func (s *Set) union(x, y int) {
	xf := s.find(x)
	yf := s.find(y)
	if xf == yf {
		return
	}
	if s.size[xf] > s.size[yf] {
		xf, yf = yf, xf
	}
	s.size[xf] += s.size[yf]
	s.parent[yf] = xf
}

func (s *Set) InSameSet(x, y int) bool {
	return s.find(x) == s.find(y)
}
func swimInWater(grid [][]int) (ret int) {
	next := [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	type Point struct {
		x, y int
	}
	N := len(grid) * len(grid)
	x := make([]Point, N)
	for i, row := range grid {
		for j, v := range row {
			x[v] = Point{x: i, y: j}
		}
	}
	set := NewSet(N)
	for cnt := 0; ; cnt++ {
		p := x[cnt]
		for i := range next {
			nx, ny := p.x+next[i][0], p.y+next[i][1]
			if nx >= 0 && ny >= 0 && nx < len(grid) && ny < len(grid) && grid[nx][ny] <= cnt {
				set.union(p.x*len(grid)+p.y, nx*len(grid)+ny)
			}
		}
		if set.InSameSet(0, N-1) {
			return cnt
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)
