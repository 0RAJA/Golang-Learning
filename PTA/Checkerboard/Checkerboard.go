package main

import (
	"fmt"
	"sort"
)

const ( //棋盘规格
	ROW = 8
	COL = 8
)

type Stack struct {
	x, y      int
	direction int         //当前方向下标
	nextIndex []NextIndex //之后的位置坐标
}

type NextIndex struct { //方便排序用
	nx, ny int
	weight int
}

var (
	next       = [8][2]int{{-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}}
	stack      [ROW*COL + 1]Stack //栈
	index      = 1                //初始栈下标
	checkBoard [ROW][COL]int      //棋盘
	weightMap  [ROW][COL]int      //权值
)

func main() {
	WeightInit() //棋盘权值初始化
	var sx, sy int
	_, err := fmt.Scanf("%d %d", &sx, &sy) //获取初始点
	if err != nil {
		fmt.Println(err)
		return
	}
	Checkerboard(sx-1, sy-1) //马踏棋盘
	Print()                  //打印
}

// Push 进栈
func Push(s Stack) {
	checkBoard[s.x][s.y] = index //标记棋盘
	stack[index] = s
	CountWeightForPoint(index) //根据权值排序
	index++
}

// Checkerboard 马踏棋盘
func Checkerboard(x, y int) {
	if x < 0 || x >= ROW || y < 0 || y >= COL {
		fmt.Println("Border Error")
		return
	}
	Push(Stack{ //入栈
		x:         x,
		y:         y,
		direction: 0,
	})
Loop:
	for index <= ROW*COL {
		p := &stack[index-1]
		Add(p.x, p.y, 1)
		for p.direction < len(p.nextIndex) {
			nx, ny := p.nextIndex[p.direction].nx, p.nextIndex[p.direction].ny
			p.direction++
			Push(Stack{ //入栈
				x:         nx,
				y:         ny,
				direction: 0,
			})
			continue Loop
		}
		index--
		checkBoard[p.x][p.y] = 0
		Add(p.x, p.y, -1)
	}
}

// IsLegal 合法性判断
func IsLegal(x, y int) bool {
	if x >= 0 && y >= 0 && x < ROW && y < COL && checkBoard[x][y] == 0 {
		return true
	}
	return false
}

// Print 打印棋盘
func Print() {
	for i := 0; i < len(checkBoard); i++ {
		for j := 0; j < len(checkBoard); j++ {
			fmt.Printf("%2d ", checkBoard[i][j])
		}
		fmt.Println()
	}
}

// WeightInit 权值初始化
func WeightInit() {
	for i, col := range weightMap {
		for j, _ := range col {
			CountWeight(i, j)
		}
	}
}

// CountWeightForPoint 计算每个点的权值
func CountWeightForPoint(index int) {
	x, y := stack[index].x, stack[index].y
	for i := 0; i < len(next); i++ {
		nx, ny := x+next[i][0], y+next[i][1]
		if IsLegal(nx, ny) {
			stack[index].nextIndex = append(stack[index].nextIndex, NextIndex{nx: nx, ny: ny, weight: CountWeight(nx, ny)})
		}
	}
	sort.Slice(stack[index].nextIndex, func(i, j int) bool {
		return stack[index].nextIndex[i].weight < stack[index].nextIndex[j].weight
	})
}

// CountWeight 计算权值
func CountWeight(x, y int) (ret int) {
	for i := 0; i < len(next); i++ {
		nx, ny := x+next[i][0], y+next[i][1]
		if IsLegal(nx, ny) {
			ret++
		}
	}
	return
}

// Add 对权值的增加和删除
func Add(x, y, t int) {
	for i := 0; i < len(next); i++ {
		nx, ny := x+next[i][0], y+next[i][1]
		if IsLegal(nx, ny) {
			weightMap[nx][ny] += t
		}
	}
}
