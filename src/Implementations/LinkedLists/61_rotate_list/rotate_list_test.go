package rotate_list

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestRotateRight(t *testing.T) {
	cases := []struct {
		input  *ListNode
		k      int
		expect *ListNode
	}{
		{SliceToList([]int{}), 1, SliceToList([]int{})},
		{SliceToList([]int{1}), 1, SliceToList([]int{1})},
		{SliceToList([]int{1, 2, 3, 4, 5}), 1, SliceToList([]int{5, 1, 2, 3, 4})},
		{SliceToList([]int{1, 2, 3, 4, 5}), 7, SliceToList([]int{4, 5, 1, 2, 3})},
		{SliceToList([]int{1, 2, 3, 4, 5}), 13, SliceToList([]int{3, 4, 5, 1, 2})},
		{SliceToList([]int{1, 2, 3, 4, 5}), 24, SliceToList([]int{2, 3, 4, 5, 1})},
		{SliceToList([]int{1, 2, 3, 4, 5}), 50, SliceToList([]int{1, 2, 3, 4, 5})},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := RotateRight(c.input, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v %v",
					ListToSlice(c.expect), ListToSlice(got), ListToSlice(c.input), c.k)
			}
		})
	}
}
