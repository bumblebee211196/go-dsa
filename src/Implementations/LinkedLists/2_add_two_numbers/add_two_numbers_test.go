package add_two_numbers

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct {
		input1 *ListNode
		input2 *ListNode
		expect *ListNode
	}{
		{SliceToList([]int{2, 4, 3}), SliceToList([]int{5, 6, 4}), SliceToList([]int{7, 0, 8})},
		{SliceToList([]int{1, 1}), SliceToList([]int{9, 9, 4}), SliceToList([]int{0, 1, 5})},
		{SliceToList([]int{5, 1, 9, 2, 6}), SliceToList([]int{4, 8, 3, 7}), SliceToList([]int{9, 9, 2, 0, 7})},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := AddTwoNumbers(c.input1, c.input2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v %v",
					ListToSlice(c.expect), ListToSlice(got), ListToSlice(c.input1), ListToSlice(c.input2))
			}
		})
	}
}
