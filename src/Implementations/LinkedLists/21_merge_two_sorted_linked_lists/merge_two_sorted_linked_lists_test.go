package _1_merge_two_sorted_linked_lists

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestMergeSortedLinkedLists(t *testing.T) {
	cases := []struct {
		input1 *ListNode
		input2 *ListNode
		expect *ListNode
	}{
		{SliceToList([]int{1, 2, 3}), SliceToList([]int{}), SliceToList([]int{1, 2, 3})},
		{SliceToList([]int{}), SliceToList([]int{4, 5, 6}), SliceToList([]int{4, 5, 6})},
		{SliceToList([]int{1, 3, 5}), SliceToList([]int{2, 4, 6}), SliceToList([]int{1, 2, 3, 4, 5, 6})},
		{SliceToList([]int{}), SliceToList([]int{}), SliceToList([]int{})},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := MergeSortedLinkedLists(c.input1, c.input2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v %v",
					ListToSlice(c.expect), ListToSlice(got), ListToSlice(c.input1), ListToSlice(c.input2))
			}
		})
	}
}
