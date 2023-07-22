package main
// import (  
//     "fmt"
// )
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)-1]
	var left, right []int

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] <= pivot {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	left = quickSort(left)
	right = quickSort(right)

	return append(append(left, pivot), right...)
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    	merged := append(nums1, nums2...)
      	mergedSort := quickSort(merged)
		var result float64
        if len(mergedSort) % 2 == 0{
            index := int(len(mergedSort)/2)
            result = (float64(mergedSort[index]) + float64(mergedSort[index-1]))/2
        } else {
			var index float64
            index = float64(len(mergedSort)/ 2)
            result = float64(mergedSort[int(index)])
			fmt.Println(index)
        }

		fmt.Println(mergedSort)
        return float64(result)
}

// func main() {
// 	nums1 := []int{1,2}
// 	nums2 := []int{3,4,5,6,8,9,10}
// 	result := findMedianSortedArrays(nums1,nums2)
// 	fmt.Println(result)
// }
