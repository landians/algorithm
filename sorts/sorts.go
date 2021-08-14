package sorts

import (
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

/* 
选择排序-将最小的数排到数组前面去, 时间复杂度O(N^2)
过程:
arr[0 ~ N-1] 范围上，找到最小值所在的位置，然后把最小值交换到 0 位置上
arr[1 ~ N-1] 范围上，找到最小值所在的位置，然后把最小值交换到 1 位置上
... 
arr[N-1 ~ N-1] 范围上，找到最小值所在的位置，然后把最小值交换到 N-1 位置上
*/
func SelectSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	for i := 0; i < len(arr)-1; i++ {
		// 假定当前数是最小的数, 记录最小数的下标
		minIndex := i
		// 从当前数的下一个数开始遍历, 选出比 minIndex 下标的值更小的数的下标
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		// 将查找出来的更小的数与当前的数进行交换, 这样就排序好了一个数
		swap(arr, i, minIndex)
	}
}

/* 
冒泡排序-将最大的数依次排到数组后面去, 时间复杂度O(N^2)
过程:
在 arr[0 ~ N-1] 范围上:
arr[0] 和 arr[1], 谁大谁来到 1 位置; arr[1] 和 arr[2], 谁大谁来到 2 位置上 ... arr[N-2] 和 arr[N-1] 谁大谁来到 N-1 位置上
在 arr[1 ~ N-2] 范围上, 重复上述过程, 但最后一步是 arr[N-3] 和 arr[N-2], 谁大谁来到 N-2 位置上
...
最后在 arr[0-1] 范围上, 重复上述过程, 但最后一步是 arr[0] 和 arr[1], 谁大谁来到 1 位置上
*/
func BubbleSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// 每遍历一次, 第 i 个数一定是最大的数了
	for i := len(arr) - 1; i > 0; i-- {
		// 相邻两个数进行比较, 若前一个数比后一个数大, 则进行交换
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}

/* 
插入排序-将最小的数依次排到数组前面去, 时间复杂度和数据初始状态有关, 如果数组数据已经有序, 则是 O(N), 最差的情况是 O(N^2), 一般选择最差的情况来估计
想让 arr[0-0] 上有序, 这个范围只有一个数, 当然是有序的
想让 arr[0-1] 上有序, 所以从 arr[1] 开始往前看, 如果 arr[1] < arr[0], 就交换, 否则什么也不做
...
想让 arr[0-i] 上有序, 所以从 arr[i] 开始往前看, 如果 arr[i] < arr[i-1], 就交换, 这样 arr[i] 这个数就在不停地
向左移动, 一直移动到左边的数字不再比自己大, 就停止移动
*/
func InsertionSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			swap(arr, j, j+1)
		}
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

// 用于生成指定大小范围和指定个数的随机数数组
func generateRandomArray(maxSize, maxValue int) []int {
	// 得到随机数的总个数
	size := rand.Intn(maxSize + 1)

	arr := make([]int, size)

	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(maxValue + 1) - rand.Intn(maxValue + 1)
	}

	return arr
}

// 用于将 arr 数组进行标准化的排序
func comparator(arr []int) {
	sort.Ints(arr)
}

// 判断数组是否有序
func ordered(arr []int) bool {
	if len(arr) == 0 {
		return true
	}

	minValue := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] < minValue {
			return false
		}
		minValue = arr[i]
	}

	return true
}
