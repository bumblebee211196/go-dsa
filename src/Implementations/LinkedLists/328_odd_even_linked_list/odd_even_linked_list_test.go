package odd_even_linked_list

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestOddEvenList(t *testing.T) {
	cases := []struct {
		input  *ListNode
		expect *ListNode
	}{
		{SliceToList([]int{1, 2, 3, 4, 5}), SliceToList([]int{1, 3, 5, 2, 4})},
		{SliceToList([]int{2, 1, 3, 5, 6, 4, 7}), SliceToList([]int{2, 3, 6, 7, 1, 5, 4})},
		{SliceToList([]int{2}), SliceToList([]int{2})},
		{SliceToList([]int{}), SliceToList([]int{})},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := OddEvenList(c.input)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v",
					ListToSlice(c.expect), ListToSlice(got), ListToSlice(c.input))
			}
		})
	}
}
