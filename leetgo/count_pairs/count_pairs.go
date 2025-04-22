package count_pairs

// countPairs counts the number of pairs (i, j) such that nums[i] == nums[j] and (i * j) % k == 0.
func countPairs(nums []int, k int) int {
    count := 0
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] == nums[j] {
                if (i * j) % k == 0 {
                    count++
                }
            }
        }
    }
    return count
}