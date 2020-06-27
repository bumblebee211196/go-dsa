package remove_nth_node_from_end_of_list

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []struct {
		input  *ListNode
		n      int
		expect *ListNode
	}{
		{SliceToList([]int{}), 0, SliceToList([]int{})},
		{SliceToList([]int{1, 2, 3, 4, 5}), 1, SliceToList([]int{1, 2, 3, 4})},
		{SliceToList([]int{1, 2, 3, 4, 5}), 2, SliceToList([]int{1, 2, 3, 5})},
		{SliceToList([]int{1, 2, 3, 4, 5}), 3, SliceToList([]int{1, 2, 4, 5})},
		{SliceToList([]int{1, 2, 3, 4, 5}), 4, SliceToList([]int{1, 3, 4, 5})},
		{SliceToList([]int{1, 2, 3, 4, 5}), 5, SliceToList([]int{2, 3, 4, 5})},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := RemoveNthFromEnd(c.input, c.n)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v %v",
					ListToSlice(c.expect), ListToSlice(got), ListToSlice(c.input), c.n)
			}
		})
	}
}
