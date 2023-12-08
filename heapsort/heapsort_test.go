package heapsort

import (
	"testing"
)

func TestCreateHeap(t *testing.T) {
	arr := []int{90, 91, 82, 77, 55, 63, 86}
	createHeap(arr, len(arr))
	t.Log(arr)
}

func TestHeapSort(t *testing.T) {
	arr := []int{2, 4, 4, 1, 9, 8, 7, 5, 0, 3}
	// 将原始数组调整为一个小顶堆
	createHeap(arr, len(arr))
	var res []int
	// 堆排序，将结果记录到结果集中
	heapSort(&res, arr, len(arr)-1)
	t.Log(res)
}

func TestRandomSlice(t *testing.T) {
	res := createRandomSlice()
	t.Log(res)
}

func TestRandomIntSlice(t *testing.T) {
	res := RandomIntSlice(10, 100)
	t.Log(res)
}

func TestRandomHeap(t *testing.T) {
	res := RandomIntSlice(10, 100)
	createHeap(res, len(res))
	t.Log(res)
}

func TestTopKNums(t *testing.T) {
	arr := []int{10, 55, 1, 2, 3, 4, 5, 6, 7, 8, 9, 78, 32}
	res := TopKNums(arr, 7)
	t.Log(res)
}
