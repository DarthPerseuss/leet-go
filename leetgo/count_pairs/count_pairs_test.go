package count_pairs

import "testing"

func TestCountPairs(t *testing.T) {
    testCases := []struct {
        nums []int
        k    int
        want int
    }{
        {[]int{1, 2, 3, 4, 5}, 2, 0},
        {[]int{1, 2, 3, 1, 1, 3}, 1, 4},
        {[]int{1, 1, 1, 1}, 1, 6},
    }

    for _, tc := range testCases {
        got := countPairs(tc.nums, tc.k)
        if got != tc.want {
            t.Errorf("countPairs(%v, %v) = %v; want %v", tc.nums, tc.k, got, tc.want)
        }
    }
}