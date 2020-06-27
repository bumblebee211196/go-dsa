package remove_duplicates_from_sorted_list

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestRemoveDuplicateNodes(t *testing.T) {
	cases := []struct {
		input  *ListNode
		expect *ListNode
	}{
		{SliceToList([]int{}), SliceToList([]int{})},
		{SliceToList([]int{1}), SliceToList([]int{1})},
		{SliceToList([]int{1, 1, 1, 2, 2, 3, 3, 3, 3}), SliceToList([]int{1, 2, 3})},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := RemoveDuplicateNodes(c.input)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v",
					ListToSlice(c.expect), ListToSlice(got), ListToSlice(c.input))
			}
		})
	}
}
