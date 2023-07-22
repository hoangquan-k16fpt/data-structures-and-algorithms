package main
// import (  
//     "fmt"
// )

func twoSum(nums []int, target int) []int {
    var n int = 1
    var k int = 0
	var arr [2]int
    for i := 0; i<len(nums)-1; i++{
        for h := n; h<len(nums); h++{
            k = nums[i]+nums[h]
            if k == target {
                arr[0] = i
                arr[1] = h
            }
        }
        n++;
    }
    return arr[:]
}

// func main() {
//     nums := []int{2,7,11,50}
//     var target int = 9
//     arr := twoSum(nums,target)
//     fmt.Println(arr)
// }
