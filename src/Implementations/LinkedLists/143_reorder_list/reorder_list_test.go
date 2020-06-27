package reorder_list

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestReorderList(t *testing.T) {
	cases := []struct {
		input  *ListNode
		expect *ListNode
	}{
		{SliceToList([]int{}), SliceToList([]int{})},
		{SliceToList([]int{1, 2, 3, 4}), SliceToList([]int{1, 4, 2, 3})},
		{SliceToList([]int{1, 2, 3, 4, 5}), SliceToList([]int{1, 5, 2, 4, 3})},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			input := CopyList(c.input)
			ReorderList(c.input)
			if !reflect.DeepEqual(c.input, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v",
					ListToSlice(c.expect), ListToSlice(c.input), ListToSlice(input))
			}
		})
	}
}
