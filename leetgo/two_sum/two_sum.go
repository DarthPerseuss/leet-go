package two_sum

// TwoSum finds two indices in nums such that their values add up to target.
func TwoSum(nums []int, target int) []int {
    for i := 0; i < len(nums)-1; i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i]+nums[j] == target {
                return []int{i, j}
            }
        }
    }
    return []int{}
}