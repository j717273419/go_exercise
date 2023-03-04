package main
import "fmt"

func main() {
	arr := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}
	fmt.Println("排序前：", arr)
	bubbleSort(arr)
	fmt.Println("排序后：", arr)
}

//实现冒泡排序
func bubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

