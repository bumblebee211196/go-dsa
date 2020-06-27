package remove_zero_sum_consecutive_nodes_from_linked_list

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestRemoveZeroSumSublists(t *testing.T) {
	cases := []struct {
		input  *ListNode
		expect *ListNode
	}{
		{SliceToList([]int{}), SliceToList([]int{})},
		{SliceToList([]int{1, 2, 3, -3, 1}), SliceToList([]int{1, 2, 1})},
		{SliceToList([]int{1, 2, -3, 3, 1}), SliceToList([]int{3, 1})},
		{SliceToList([]int{1, 2, 3, -3, -2}), SliceToList([]int{1})},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := RemoveZeroSumSublists(c.input)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v",
					ListToSlice(c.expect), ListToSlice(got), ListToSlice(c.input))
			}
		})
	}
}
