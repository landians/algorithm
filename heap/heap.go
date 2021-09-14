package heap

import (
	"math/rand"
	"sort"
)

// heapInsert 当新的节点插入到堆时，将其放到合适的位置， 时间复杂度：O(logN)， 按照高度走
func heapInsert(values []int, i int) {
	// 当前节点的值比其父节点的值大时，则交换两个节点的位置
	for values[i] > values[(i-1)/2] {
		values[i], values[(i-1)/2] = values[(i-1)/2], values[i]
		i = (i - 1) / 2
	}
}

// heapIfy 用于调整数组重新为堆， 时间复杂度：O(logN)， 按照高度走
func heapIfy(values []int, i int, heapSize int) {
	left := 2*i + 1

	// 因为 right 的坐标是大于 left, 所以如果 left >= heapSize，那说明当前节点肯定没有子节点了
	for left < heapSize {

		largest := 0

		// left + 1 就是  right 了，这里是为了找到左右中最大的那一个
		if left+1 < heapSize && values[left+1] > values[left] {
			largest = left + 1
		} else {
			largest = left
		}

		// 再与父节点进行比较得到最终的最大值的下标
		if values[largest] < values[i] {
			largest = i
		}

		// 因为左右子节点的值并没有大于父节点，所以并不需要继续操作了
		if largest == i {
			break
		}

		values[largest], values[i] = values[i], values[largest]
		i = largest
		left = 2*i + 1
	}
}

// heapSort 堆排序，时间复杂度: O(NlogN), 空间复杂度O(1)
func heapSort(values []int) {
	if len(values) < 2 {
		return
	}

	// heapInsert 数组元素建立堆
	//for i := 0; i < len(values); i++ { // O(N)
	//	heapInsert(values, i) // O(logN)
	//}

	// 这种方式的建堆更快
	for i := (len(values)/2) -1; i >=0; i-- { // O(N)
		heapIfy(values, i, len(values)) // O(logN)
	}

	heapSize := len(values)

	// 将堆顶点的下标的元素和堆尾的下标的元素进行交换，heapSize --
	heapSize--
	values[0], values[heapSize] = values[heapSize], values[0]

	// 从堆顶点的下标开始进行 heapIfy 操作，调整堆, 知道 heapSize == 0
	for heapSize > 0 { // O(N)
		heapIfy(values, 0, heapSize) // O(logN)
		heapSize--
		values[0], values[heapSize] = values[heapSize], values[0]
	}
}

// 用于生成指定大小范围和指定个数的随机数数组
func generateRandomArray(maxSize, maxValue int) []int {
	// 得到随机数的总个数
	size := rand.Intn(maxSize + 1)

	arr := make([]int, size)

	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(maxValue+1) - rand.Intn(maxValue+1)
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
