package binarysearch

import "sort"

func SearchInts(list []int, key int) int {
	sort.Ints(list)
    left, right := 0, len(list) - 1
    var mid int
    for left <= right {
        mid = (left + right) / 2
        if list[mid] == key {
            return mid
        }
        if list[mid] < key {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
