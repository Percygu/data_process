package heapsort

import (
	"fmt"
	"math/rand"
	"time"
)

/*
从数组中下标为i的节点开始调整，使得调整完之后数组继续维持堆的属性,这里以小顶堆为例
@param arr 数组
@param n 数组长度
@param i 待维护节点的下标
*/
func heapify(arr []int, n, i int) {
	largest, l, r := i, 2*i+1, 2*i+2
	// 找到下标为i的节点以及他的左右孩子中值最大的节点的下标，记录到largest中
	if l < n && arr[largest] > arr[l] {
		largest = l
	}
	if r < n && arr[largest] > arr[r] {
		largest = r
	}
	// 值最大的节点的下标不是i，说明值最大的节点在其左右孩子中，交换这个值最大的孩子节点和其父节点
	if i != largest {
		arr[largest], arr[i] = arr[i], arr[largest]
		heapify(arr, n, largest)
	}
}

// heapSort 这里演示小顶堆，从小到大排序
func createHeap(arr []int, n int) {
	// 完全二叉树，从最后一个有孩子的节点开始向下调整即可
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

func heapSort(res *[]int, arr []int, n int) {
	*res = append(*res, arr[0])
	arr[0], arr[n] = arr[n], arr[0]
	n--
	if n < 0 {
		return
	}
	heapify(arr, n+1, 0) // 调整堆
	heapSort(res, arr, n)
}

// Random 得到一个随机数
func Random(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if max < 1 {
		return r.Int()
	} else {
		return r.Intn(max)
	}
}

func RandomIntSlice(length int, maxValue int) []int {
	var res []int
	for i := 0; i < length; i++ {
		num := Random(maxValue)
		res = append(res, num)
	}
	fmt.Printf("res = %+v\n", res)
	return res
}

func createRandomSlice() []int {
	var res []int
	for i := 0; i < 20; i++ {
		res = append(res, Random(10000))
	}
	return res
}

func TopKNums(arr []int, k int) []int {
	if k >= len(arr) {
		return arr
	}
	nums := arr[:k]
	createHeap(nums, len(nums)) // 将nums数组调整为小顶堆
	for _, num := range arr[k:] {
		if num <= nums[0] {
			continue
		}
		nums[0] = num
		heapify(nums, k, 0)
	}
	return nums
}

func main() {
	arr := []int{2, 4, 6, 1, 9, 8, 7, 5, 0, 3}
	createHeap(arr, len(arr))
	fmt.Println(arr)
}
