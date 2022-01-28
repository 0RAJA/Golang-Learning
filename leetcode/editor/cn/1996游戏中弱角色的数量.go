//你正在参加一个多角色游戏，每个角色都有两个主要属性：攻击 和 防御 。给你一个二维整数数组 properties ，其中 properties[i] = [
//attacki, defensei] 表示游戏中第 i 个角色的属性。
//
// 如果存在一个其他角色的攻击和防御等级 都严格高于 该角色的攻击和防御等级，则认为该角色为 弱角色 。更正式地，如果认为角色 i 弱于 存在的另一个角色
//j ，那么 attackj > attacki 且 defensej > defensei 。
//
// 返回 弱角色 的数量。
//
//
//
// 示例 1：
//
//
//输入：properties = [[5,5],[6,3],[3,6]]
//输出：0
//解释：不存在攻击和防御都严格高于其他角色的角色。
//
//
// 示例 2：
//
//
//输入：properties = [[2,2],[3,3]]
//输出：1
//解释：第一个角色是弱角色，因为第二个角色的攻击和防御严格大于该角色。
//
//
// 示例 3：
//
//
//输入：properties = [[1,5],[10,4],[4,3]]
//输出：1
//解释：第三个角色是弱角色，因为第二个角色的攻击和防御严格大于该角色。
//
//
//
//
// 提示：
//
//
// 2 <= properties.length <= 10⁵
// properties[i].length == 2
// 1 <= attacki, defensei <= 10⁵
//
// Related Topics 栈 贪心 数组 排序 单调栈 👍 101 👎 0

package main

import "sort"

func main() {

}

//leetcode submit region begin(Prohibit modification and deletion)
func numberOfWeakCharacters(properties [][]int) (ret int) {
	sort.Slice(properties, func(i, j int) bool {
		return properties[i][0] > properties[j][0] || properties[i][0] == properties[j][0] && properties[i][1] < properties[j][1]
	})
	maxDefense := 0
	for i := range properties {
		if properties[i][1] < maxDefense {
			ret++
		} else {
			maxDefense = properties[i][1]
		}
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
