package reverse_linked_list_2

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestReverseBetween(t *testing.T) {
	cases := []struct {
		input  *ListNode
		m, n   int
		expect *ListNode
	}{
		{SliceToList([]int{}), 0, 0, SliceToList([]int{})},
		{SliceToList([]int{1}), 1, 1, SliceToList([]int{1})},
		{SliceToList([]int{1, 2, 3, 4, 5}), 2, 4, SliceToList([]int{1, 4, 3, 2, 5})},
		{SliceToList([]int{1, 2, 3, 4, 5}), 1, 4, SliceToList([]int{4, 3, 2, 1, 5})},
		{SliceToList([]int{1, 2, 3, 4, 5}), 1, 5, SliceToList([]int{5, 4, 3, 2, 1})},
		{SliceToList([]int{1, 2, 3, 4, 5}), 3, 3, SliceToList([]int{1, 2, 3, 4, 5})},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := ReverseBetween(c.input, c.m, c.n)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v %v, %v",
					ListToSlice(c.expect), ListToSlice(got), ListToSlice(c.input), c.m, c.n)
			}
		})
	}
}
