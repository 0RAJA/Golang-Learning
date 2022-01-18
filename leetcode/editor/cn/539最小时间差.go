//给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。
//
//
//
// 示例 1：
//
//
//输入：timePoints = ["23:59","00:00"]
//输出：1
//
//
// 示例 2：
//
//
//输入：timePoints = ["00:00","23:59","00:00"]
//输出：0
//
//
//
//
// 提示：
//
//
// 2 <= timePoints <= 2 * 10⁴
// timePoints[i] 格式为 "HH:MM"
//
// Related Topics 数组 数学 字符串 排序 👍 140 👎 0

package main

import (
	"fmt"
	"math"
	"sort"
	time2 "time"
)

func main() {
	fmt.Println(findMinDifference([]string{"07:56", "19:58", "19:12", "01:59", "04:27"}))
}

//leetcode submit region begin(Prohibit modification and deletion)
func findMinDifference(timePoints []string) (ret int) {
	ret = math.MaxInt32
	abs := func(duration float64) float64 {
		if duration < 0 {
			return -duration
		}
		return duration
	}
	const Format = "15:04"
	pro, _ := time2.Parse(Format, "00:00")
	mid, _ := time2.Parse(Format, "12:00")
	post, _ := time2.Parse(Format, "23:59")
	sub := func(t time2.Time) float64 {
		if t.Before(mid) {
			return t.Sub(pro).Minutes()
		}
		return post.Sub(t).Minutes()
	}
	sort.Strings(timePoints)
	for i := 0; i < len(timePoints); i++ {
		sx, sy := i, i+1
		if i+1 == len(timePoints) {
			sy = 0
		}
		t1, _ := time2.Parse(Format, timePoints[sx])
		t2, _ := time2.Parse(Format, timePoints[sy])
		tmp := abs(t1.Sub(t2).Minutes())
		x1 := sub(t1)
		x2 := sub(t2)
		if tmp > x1+x2+1 {
			tmp = x1 + x2 + 1
		}
		if int(tmp) < ret {
			ret = int(tmp)
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
