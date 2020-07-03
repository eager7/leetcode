package twoSum

/*
自己实现的第一版，考虑了多个相同元素必须按照顺序返回的情况。
先把数据按照值-下标的模式存储到map，多个下标用切片存储。
然后遍历nums，不去遍历map是因为map随机，不能按照顺序返回，遍历nums时找出和target之差，去map里查看是否有这个值，
如果有则返回，这里注意剔除6=3+3的情况。

作者：eager7
链接：https://leetcode-cn.com/problems/two-sum/solution/golangjie-ti-by-eager7/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

func twoSum(nums []int, target int) []int {
	var m = make(map[int][]int)
	for index, n := range nums {
		m[n] = append(m[n], index)
	}
	for j, k := range nums {
		if indexs, ok := m[target-k]; ok {
			for _, v := range indexs { //有多个匹配项，找出不等的返回
				if v == j { //是自身
					continue
				}
				return []int{j, v}
			}
		}
	}
	return nil
}
