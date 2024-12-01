package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {

	nums1 := []int{3, 4, 2, 1, 3, 3}
	nums2 := []int{4, 3, 5, 3, 9, 3}

	sort.Ints(nums1)
	sort.Ints(nums2)

	var result float64
	for i := 0; i < len(nums1); i++ {
		result += (math.Abs(float64(nums1[i]) - float64(nums2[i])))
	}

	fmt.Println(result)

	//TODO: read data from csv file

}
