package main

import (
	"fmt"
	"math/rand"	
	"time"
)

func main()  {
	arr := []int{}	//随机产生10个0~100中的数
	rand.Seed(time.Now().UnixNano())	//真随机数
    for i := 0; i < 10; i++ {
        arr = append(arr, rand.Intn(100))
	}
	fmt.Println("Input:", arr)
    quickSort(arr, 0, len(arr)-1)
    fmt.Println("Result:", arr)
}

func quickSort(arr []int, start, end int)  {
	if (start < end) {
        // 以数组的第一个值为主元（参照值）
        pivot := arr[start]
        i := start
        j := end
        
        for i < j {
            // 从右边找到第一个比pivot小的值
            for i < j && arr[j] >= pivot {
                j--
            }
            if i < j {
                arr[i] = arr[j]
                i++
            }
            // 从左边找到第一个比pivot大的值
            for i < j && arr[i] <= pivot {
                i++
            }
            if i < j {
                arr[j] = arr[i]
                j--
            }
        }
 
		arr[i] = pivot
		fmt.Println("QuickSorting:", arr)
        // 分两组再递归排序
        quickSort(arr, start, i - 1)
        quickSort(arr, i + 1, end)
    }
}