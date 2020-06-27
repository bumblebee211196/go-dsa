package linked_list_components

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestNumComponents(t *testing.T) {
	cases := []struct {
		input  *ListNode
		G      []int
		expect int
	}{
		{SliceToList([]int{1, 2, 3, 4, 5, 6}), []int{9, 8, 7}, 0},
		{SliceToList([]int{0, 1, 2, 3}), []int{0, 1, 3}, 2},
		{SliceToList([]int{0, 1, 2, 3, 4}), []int{0, 3, 1, 4}, 2},
		{SliceToList([]int{0, 1, 2, 3, 4}), []int{4, 3, 2, 1, 0}, 1},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := NumComponents(c.input, c.G)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v %v",
					c.expect, got, ListToSlice(c.input), c.G)
			}
		})
	}
}
