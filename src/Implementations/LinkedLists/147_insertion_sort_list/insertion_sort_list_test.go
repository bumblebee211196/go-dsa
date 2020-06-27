package insertion_sort_list

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestInsertionSortList(t *testing.T) {
	cases := []struct {
		input  *ListNode
		expect *ListNode
	}{
		{SliceToList([]int{}), SliceToList([]int{})},
		{SliceToList([]int{4, 2, 1, 3}), SliceToList([]int{1, 2, 3, 4})},
		{SliceToList([]int{-1, 5, 3, 4, 0}), SliceToList([]int{-1, 0, 3, 4, 5})},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := InsertionSortList(c.input)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v",
					ListToSlice(c.expect), ListToSlice(got), ListToSlice(c.input))
			}
		})
	}
}
