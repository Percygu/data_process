package mergesort

import "testing"

func TestMergesort(t *testing.T) {
	t.Log(RandomDataSlice(10, 20, 1000))
}

func TestHeapSort(t *testing.T) {
	nums := []Data{}
	for j := 0; j < 20; j++ {
		num := Random(10000)
		nums = append(nums, Data{
			Value: num,
			Index: 1,
		})
	}
	t.Log(nums)
	// 将原始数组调整为一个小顶堆
	createHeap(nums, len(nums))
	var res []int
	heapSort(&res, nums, len(nums)-1)
	t.Log(res)
}

func TestMultiMergeSort(t *testing.T) {
	//nums := RandomIntSlice(5, 10, 100)
	nums := [][]int{{3, 4, 13, 15, 23}, {7, 26, 39, 39, 75}, {27, 63, 71, 88, 93}}
	t.Log(nums)
	res := multiMergeSort(nums)
	t.Log(res)
}

func TestRandomIntSlice(t *testing.T) {
	res := RandomIntSlice(7, 5, 100)
	t.Log(res)
}
