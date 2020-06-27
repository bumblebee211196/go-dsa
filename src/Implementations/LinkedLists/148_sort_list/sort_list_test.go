package sort_list

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestSortList(t *testing.T) {
	cases := []struct {
		input  *ListNode
		expect *ListNode
	}{
		{SliceToList([]int{}), SliceToList([]int{})},
		{SliceToList([]int{1}), SliceToList([]int{1})},
		{SliceToList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}), SliceToList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})},
		{SliceToList([]int{1, 7, 3, 9, 4, 2, 8, 1, 4, 2, 8}), SliceToList([]int{1, 1, 2, 2, 3, 4, 4, 7, 8, 8, 9})},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := SortList(c.input)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v",
					c.expect, got, c.input)
			}
		})
	}
}
