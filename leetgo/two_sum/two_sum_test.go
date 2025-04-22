package two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
    testCases := []struct {
        nums   []int
        target int
        want   []int
    }{
        {[]int{2, 7, 11, 15}, 9, []int{0, 1}},
        {[]int{3, 2, 4}, 6, []int{1, 2}},
        {[]int{0, 4, 3, 0}, 0, []int{0, 3}},
        {[]int{-1, -2, -3, -4, -5}, -8, []int{2, 4}},
    }

    for _, tc := range testCases {
        got := TwoSum(tc.nums, tc.target)
        if !reflect.DeepEqual(got, tc.want) {
            t.Errorf("TwoSum(%v, %v) = %v; want %v", tc.nums, tc.target, got, tc.want)
        }
    }
}