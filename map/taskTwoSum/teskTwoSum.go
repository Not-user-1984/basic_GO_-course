package tasktwosum

func TwoSum(target int ,nums[]int) []int{

	numIndexMap := make(map[int]int)

	for i , num := range nums{
		complate := target - num
		if index, found := numIndexMap[complate]; found {
			return []int{index, i }
		}
		numIndexMap[num] = i
	}
	return []int{}
}