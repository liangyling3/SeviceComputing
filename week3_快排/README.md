## 快速排序算法简介
快速排序（Quicksort）是对冒泡排序的一种改进。

快速排序由C. A. R. Hoare在1960年提出。它的基本思想是：通过一趟排序将要排序的数据分割成独立的两部分，其中一部分的所有数据都比另外一部分的所有数据都要小，然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，以此达到整个数据变成有序序列。
### 基本思想
从这个数列里找一个数作为基准点（支点）跟其它的数进行对比，小于或等于这个支点数的都放到左边，大的放到右边。完成一次排序，然后依次递归，当起始下标和末尾下标相遇是，排序结束，即整列数已经排列好了顺序。

### 步骤
1. 从数列中选取一个数作为基准点。
2. 从右向左寻找小于基准点的值，并放到左边（索引j）；从左向右寻找大于基准点的值，并放到右边（索引i）。交替进行，直到 i=j。
3. 将基准点放入最后一个数组空位，此时基准点左侧的值全部小于基准点右侧的值。
4. 递归，将基准点分割出的两部分再分别进行排序。

### 代码
```go
package main

import (
	"fmt"
	"math/rand"	
	"time"
)

// 主函数
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

//快排函数
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
```

### 测试结果
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190915144605443.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2xpYW5neWxpbmcz,size_16,color_FFFFFF,t_70)
![在这里插入图片描述](https://img-blog.csdnimg.cn/20190915144623803.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2xpYW5neWxpbmcz,size_16,color_FFFFFF,t_70)
