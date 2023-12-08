package mergesort

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Data struct {
	Value int // 值
	Index int // 数组索引号，标识来自于那个数组
}

func (d *Data) GetValue() int {
	return d.Value
}

func (d *Data) GetIndex() int {
	return d.Index
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

func RandomDataSlice(sliceSize, length int, maxValue int) [][]Data {
	var res [][]Data
	for i := 0; i < sliceSize; i++ {
		nums := []Data{}
		for j := 0; j < length; j++ {
			num := Random(maxValue)
			nums = append(nums, Data{
				Value: num,
				Index: i,
			})
		}
		res = append(res, nums)
	}
	fmt.Printf("res = %+v\n", res)
	return res
}

func RandomIntSlice(sliceSize, length int, maxValue int) [][]int {
	var res [][]int
	for i := 0; i < sliceSize; i++ {
		nums := []int{}
		for j := 0; j < length; j++ {
			num := Random(maxValue)
			nums = append(nums, num)
		}
		sort.Ints(nums) // 堆nums进行排序
		res = append(res, nums)
	}
	fmt.Printf("res = %+v\n", res)
	return res
}

/*
从数组中下标为i的节点开始调整，使得调整完之后数组继续维持堆的属性,这里以小顶堆为例
@param arr 数组
@param n 数组长度
@param i 待维护节点的下标
*/
func heapify(arr []Data, n, i int) {
	largest, l, r := i, 2*i+1, 2*i+2
	// 找到下标为i的节点以及他的左右孩子中值最大的节点的下标，记录到largest中
	if l < n && arr[largest].Value > arr[l].Value {
		largest = l
	}
	if r < n && arr[largest].Value > arr[r].Value {
		largest = r
	}
	// 值最大的节点的下标不是i，说明值最大的节点在其左右孩子中，交换这个值最大的孩子节点和其父节点
	if i != largest {
		arr[largest], arr[i] = arr[i], arr[largest]
		heapify(arr, n, largest)
	}
}

// createHeap 创建小顶堆
func createHeap(arr []Data, n int) {
	// 完全二叉树，从最后一个有孩子的节点开始向下调整即可
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
}

// heapSort 小顶堆排序,maxIndex为数组的最大索引编号
func heapSort(res *[]int, arr []Data, maxIndex int) {
	*res = append(*res, arr[0].Value)
	arr[0], arr[maxIndex] = arr[maxIndex], arr[0]
	maxIndex--
	if maxIndex < 0 {
		return
	}
	heapify(arr, maxIndex+1, 0) // 调整堆，maxIndex+1为数组的长度
	heapSort(res, arr, maxIndex)
}

// multiMergeSort 多路归并
func multiMergeSort(numList [][]int) []int {
	var res []int   // 排序之后的结果集
	var nums []Data // num标识堆
	n := len(numList)
	for i := 0; i < len(numList); i++ {
		data := Data{
			Value: numList[i][0],
			Index: i,
		}
		nums = append(nums, data)
	}
	//以每个数组的第0号元素初始化Data类型的小顶堆
	createHeap(nums, len(numList))
	fmt.Printf("nums=%v\n", nums)
	signs := make([]int, n)
	// 标志位，有len(numList)个数组，所以有len(numList)个标识位，每个标志位从1开始，因为第0号元素已经用来构建堆了，signs[i]用于表示当前数组下一个要读取的位置
	for i := 0; i < len(numList); i++ {
		signs[i] = 1
	}

	var (
		currentDeleteData     Data // 当前从堆顶移除的元素
		currentInsertData     Data // 当前插入到堆顶的元素
		arrIndexForDeleteData int  // 从堆顶移除的元素所在数组索引
		nextValueIndexInArray int  // 下一个待插入堆顶的元素在数组的索引
	)
	//  不停的读取数组中的元素，直到所有的数组中都没有元素了
	for {
		var ArrayEnded = 0 // 维护有多少个数组已经读取完了
		for i := 0; i < n; i++ {
			if signs[i] >= len(numList[i]) {
				ArrayEnded++
			}
		}
		if ArrayEnded == n {
			heapSort(&res, nums, len(nums)-1)
			break
		}
		currentDeleteData = nums[0] // 删除堆顶元素，因为堆顶元素最小
		fmt.Printf("currentDeleteData = %v\n", currentDeleteData)
		res = append(res, currentDeleteData.Value)           // 将堆顶元素记录到结果集中
		arrIndexForDeleteData = currentDeleteData.GetIndex() // 删除的堆顶元素所在的数组索引号
		nextValueIndexInArray = signs[arrIndexForDeleteData] // 下一个待插入堆顶的元素在数组的索引位置
		// 当前这个数组还没有读取完
		if nextValueIndexInArray < len(numList[arrIndexForDeleteData]) {
			currentInsertData = Data{
				Value: numList[arrIndexForDeleteData][nextValueIndexInArray],
				Index: arrIndexForDeleteData,
			}
			nums[0] = currentInsertData
			// 重新调整堆,置换一次堆顶，调整一次堆
			heapify(nums, len(nums), 0)
			// 本数组种下一个带插入到堆顶的元素位置后移一个单位
			signs[arrIndexForDeleteData]++
		} else {
			// 当前数组已经读取完了,说明这个数组没有任何元素了，查找下一个有值数组的最小元素
			for {
				arrIndexForDeleteData = (arrIndexForDeleteData + 1) % n // 获取到有值数组的索引号
				nextValueIndexInArray = signs[arrIndexForDeleteData]    // 下一个待插入堆顶的元素在这个数组的索引位置
				if nextValueIndexInArray >= len(numList[arrIndexForDeleteData]) {
					continue
				}
				// 找到了元素，插入到堆顶
				currentInsertData = Data{
					Value: numList[arrIndexForDeleteData][nextValueIndexInArray],
					Index: arrIndexForDeleteData,
				}
				nums[0] = currentInsertData
				// 重新调整堆,置换一次堆顶，调整一次堆
				heapify(nums, len(nums), 0)
				// 本数组种下一个带插入到堆顶的元素位置后移一个单位
				signs[arrIndexForDeleteData]++
				break
			}
		}
	}
	return res
}
