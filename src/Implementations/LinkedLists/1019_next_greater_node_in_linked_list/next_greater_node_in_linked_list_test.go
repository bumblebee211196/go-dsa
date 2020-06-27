package next_greater_node_in_linked_list

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestNextLargerNodes(t *testing.T) {
	cases := []struct {
		input  *ListNode
		expect []int
	}{
		{SliceToList([]int{}), []int{}},
		{SliceToList([]int{2, 1, 5}), []int{5, 5, 0}},
		{SliceToList([]int{2, 7, 4, 3, 5}), []int{7, 0, 5, 5, 0}},
		{SliceToList([]int{1, 7, 5, 1, 9, 2, 5, 1}), []int{7, 9, 9, 9, 0, 5, 0, 0}},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := NextLargerNodes(c.input)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v",
					c.expect, got, ListToSlice(c.input))
			}
		})
	}
}
