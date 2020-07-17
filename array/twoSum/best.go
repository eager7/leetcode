package twoSum

func best(nums []int, target int) []int {
	var m = make(map[int][]int)
	for index, n := range nums {
		if indexs, ok := m[target-n]; ok {
			for _, v := range indexs { //有多个匹配项，找出不等的返回
				if v == index { //是自身
					continue
				}
				return []int{v, index} //这里先存的v因此需要先返回v
			}
		}
		m[n] = append(m[n], index)
	}
	return nil
}
