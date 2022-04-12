package main

import "fmt"

func main() {
	data := []int{1, 3, 9, 11, 13, 22, 22, 22, 22, 27, 56, 73, 88, 88, 99, 143, 175, 176, 201, 300}

	l, r := RepeatValueBinarySearch(data, 88)
	fmt.Printf("left index: %v\nright index: %v\n", l, r)

}

func BinarySearch(data []int, key int) (index int) {
	fmt.Printf("BinarySearch sorted data: %v, target data: %v\n", data, key)
	low := 0
	upper := len(data) - 1

	// 閉區間
	// 等於的時候要在比較一次該數值
	for low <= upper {
		// 無條件捨去，只取商數
		// low: 0, upper: 11, targetIndex會等於5
		targetIndex := (low + upper) / 2
		fmt.Printf("low: %v, upper: %v, targetIndex: %v\n", low, upper, targetIndex)

		if low == upper {
			v := data[targetIndex]
			if v == key {
				return targetIndex
			}
			// 找不到正確資料，回傳相差最小的資料
			// 比較此時的target index的值，如果比target小，那就要找target index + 1 中的值再比一次，比較大則相反
			switch {
			case v == key:
				return targetIndex
			case v < key:
				// avoid out of slice index
				if targetIndex+1 > len(data)-1 {
					return targetIndex
				}
				// 如果 v < key，就要再找一個比key大的，形成 v < key < v2，再比大小
				v2 := data[targetIndex+1]
				if key-v > v2-key {
					// v2和key相差的值最小，回傳v2的index值
					return targetIndex + 1
				} else {
					return targetIndex
				}
			case v > key:
				// avoid out of slice index
				if targetIndex-1 < 0 {
					return targetIndex
				}
				// 如果 v > key，就要再找一個比key小的，形成 v2 < key < v，再比大小
				v2 := data[targetIndex-1]
				if v-key > key-v2 {
					// v2和key相差的值最小，回傳v2的index值
					return targetIndex - 1
				} else {
					return targetIndex
				}
			}
		}

		// targetIndex +- 1，target Index的值已經比較過，就直接排除掉。
		switch {
		case data[targetIndex] == key:
			// 重複值 todo
			return targetIndex
		case data[targetIndex] > key:
			upper = targetIndex - 1
			fmt.Printf("data[targetIndex] %v > targetData %v , upper: %v\n", data[targetIndex], key, upper)
		case data[targetIndex] < key:
			low = targetIndex + 1
			fmt.Printf("data[targetIndex] %v < targetData %v , low: %v\n", data[targetIndex], key, low)
		}
		fmt.Printf("for loop process over low: %v, upper: %v\n\n", low, upper)
	}
	// fmt.Printf("final low: %v, upper: %v\n", low, upper)
	// fmt.Print("data not found\n")

	// return low
	return -1
}

func RepeatValueBinarySearch(data []int, key int) (leftIndex int, rightIndex int) {
	low := 0
	upper := len(data) - 1
	fmt.Printf("BinarySearch sorted data: %v, target data: %v, len: %v\n", data, key, len(data))

	// 先找到重複的target value最小的index
	for low <= upper {
		mi := (low + upper) / 2
		switch {
		case data[mi] < key:
			low = mi + 1
		// 等於的時候也要往左移動upper index，因為最終要找的是>=target data中最小的index，直到不等於key值，再加一回去就是index值最小的重複值
		case data[mi] >= key:
			upper = mi - 1
		}
	}
	// leftIndex = upper + 1
	leftIndex = low

	low2 := 0
	upper2 := len(data) - 1

	// 找到不等於重複的target value最小的index
	for low2 <= upper2 {
		mi := (low2 + upper2) / 2
		switch {
		// 等於的時候也要往右移動index(加一)，直到不等於key值，再減一回去就是index值最小的重複值
		case data[mi] <= key:
			low2 = mi + 1
		case data[mi] > key:
			upper2 = mi - 1
		}
	}
	// rightIndex = low2 - 1
	rightIndex = upper2

	// 再找到大於target value最小的index
	return
}
