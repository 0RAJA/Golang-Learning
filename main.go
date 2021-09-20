package main

import "fmt"

var Map [9][9]int  //存储路径
var book [9][9]int //判断是否走过
var next = [8][2]int{{-2, 1}, {-1, 2}, {1, 2}, {2, 1}, {2, -1}, {1, -2}, {-1, -2}, {-2, -1}}
var x = 1
var m, n int
var IsOK bool

func main() {
	fmt.Scan(&m, &n)
	Map[m][n] = x
	book[m][n] = 1
	x++
	CheckerBored(m, n)
	print()
}
func CheckerBored(m, n int) {
	for k := 0; k < 8; k++ {
		tx := m + next[k][0]
		ty := n + next[k][1]
		if tx > 8 || tx < 1 || ty > 8 || ty < 1 { //如果越界就不考虑
			continue
		}
		if book[tx][ty] == 0 { //如果这个地方没有走到，就走这个地方
			book[tx][ty] = 1 //每次走到一个没走到的地方就标记
			Map[tx][ty] = x  //记录路径
			x++
			CheckerBored(tx, ty) //递归搜索
			if IsOK {
				return
			}
			if x > 64 {
				IsOK = true
			}
			book[tx][ty] = 0 //回溯的时候取消标记
			Map[tx][ty] = 0
			x--
		}
	}
}
func print() {
	for i := 1; i <= 8; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Printf("%d ", Map[i][j])
		}
		fmt.Printf("\n")
	}
}
