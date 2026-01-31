package armstrong

import "math"  

func IsNumber(n int) bool {
	nums := make([]int, 0)
    buf := n
    sum := 0
    for buf > 0 {
        nums = append(nums, buf % 10)
        buf /= 10
    }
    for i := 0; i < len(nums); i++ {
        sum += int(math.Pow(float64(nums[i]), float64(len(nums))))
    }
    return sum == n
}
