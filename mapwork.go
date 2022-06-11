func findLeastNumOfUniqueInts(arr []int, k int) int {
	// Map用来统计数字出现的频率
	// Arr为数字出现频率数组
	arrMap, arrNum := make(map[int]int), make([]int, 0)
	for i := 0; i < len(arr); i ++ {
		arrMap[arr[i]] ++
	}

	for _, num := range arrMap {
		arrNum = append(arrNum, num)
	}
	// 排序
	sort.Ints(arrNum)

	cnt := 0
	for i, v := range arrNum{
		if k > v {	// 若当前k大于该频率,则直接减
			k -= v
		} else if k == v {	//	若等于,则最少数目为接下来数组的长度
			cnt = len(arrNum) - i - 1
			break
		} else { //	若小于,则最少数目为当前数之后数组的长度
			cnt = len(arrNum) - i
			break
		}
	}

	return cnt

}