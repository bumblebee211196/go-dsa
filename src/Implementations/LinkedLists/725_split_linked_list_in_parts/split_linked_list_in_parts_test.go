package split_linked_list_in_parts

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestSplitListToParts(t *testing.T) {
	cases := []struct {
		input1 *ListNode
		input2 int
		expect []*ListNode
	}{
		{SliceToList([]int{1, 2, 3}), 5,
			[]*ListNode{
				SliceToList([]int{1}),
				SliceToList([]int{2}),
				SliceToList([]int{3}),
				SliceToList([]int{}),
				SliceToList([]int{}),
			},
		},
		{SliceToList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}), 3,
			[]*ListNode{
				SliceToList([]int{1, 2, 3, 4}),
				SliceToList([]int{5, 6, 7}),
				SliceToList([]int{8, 9, 10}),
			},
		},
		{SliceToList([]int{1, 2, 3, 4, 5, 6, 7, 8}), 5,
			[]*ListNode{
				SliceToList([]int{1, 2}),
				SliceToList([]int{3, 4}),
				SliceToList([]int{5, 6}),
				SliceToList([]int{7}),
				SliceToList([]int{8}),
			},
		},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := SplitListToParts(c.input1, c.input2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v %v",
					c.expect, got, ListToSlice(c.input1), c.input2)
			}
		})
	}
}
