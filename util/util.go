package util

import (
	"sort"
)

func RemoveElements(slice []interface{}, indices []int) {
	// 将要删除的索引从小到大排序
	sort.Ints(indices)

	// 遍历要删除的索引，每次删除一个元素
	for i := len(indices) - 1; i >= 0; i-- {
		index := indices[i]
		slice = append(slice[:index], slice[index+1:]...)
	}

}
