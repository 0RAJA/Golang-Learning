//给你一个整数 n ，请你在无限的整数序列 [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...] 中找出并返回第 n 位数字。
//
//
//
// 示例 1：
//
//
//输入：n = 3
//输出：3
//
//
// 示例 2：
//
//
//输入：n = 11
//输出：0
//解释：第 11 位数字在序列 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... 里是 0 ，它是 10 的一部分。
//
//
//
//
// 提示：
//
//
// 1 <= n <= 2³¹ - 1
//
// Related Topics 数学 二分查找 👍 233 👎 0

package main

import "fmt"

func main() {
	fmt.Println(findNthDigit2(11))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findNthDigit(n int) (ret int) {
	/*
		index  min 	 	max  	数字个数		总位数
		1	   1		9	 	9			1*9
		2      10   	99   	90			2*90
		3      100   	999   	900			3*900
		4      1000   	9999   	9000		4*9000
		...
		所以计算一个数num对应的值ret,先算num所属index的值,即num是在哪个区间范围[min,max]内,然后再计算num在index区间中的哪个数n上,再计算是n上的哪一位ret
		举例: num = 11
		计算得: index = 2 即num在(index = 2)区间内,区间第一个数字为10,而11相对于第一个数的位数为x = num-(1*9) = 2,则数字偏移量pos = ((x-1)/index) = 0
				所以num对应的数字ret是第index区间的数字 n = min + pos = 10 上的第(从左往右) n-index*pos 位.
	*/
	var index = 1
	m := 9
	for index <= 10 { //找到区间
		if n <= m*index {
			break
		}
		n -= m * index
		m *= 10
		index++
	}
	pos := (n - 1) / index //计算相对此区间第一个数的偏移量
	m = m/9 + pos          //通过偏移量找到这个数
	n -= index * pos       //计算在这个数的第几位
	for i := 0; i <= index-n; i++ {
		ret = m % 10
		m /= 10
	}
	return
}
func findNthDigit2(n int) (ret int) {
	idx := 1
	m := 9
	for idx <= 10 {
		if idx*m >= n {
			break
		}
		n -= idx * m
		idx++
		m *= 10
	}
	pos := (n - 1) / idx //计算相对此区间第一个数的偏移量
	m = m/9 + pos        //通过偏移量找到这个数
	n -= idx * pos       //计算在这个数的第几位
	for i := 0; i <= idx-n; i++ {
		ret = m % 10
		m /= 10
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
