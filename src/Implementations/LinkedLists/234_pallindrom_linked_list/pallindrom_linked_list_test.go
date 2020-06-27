package pallindrom_linked_list

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		input  *ListNode
		expect bool
	}{
		{SliceToList([]int{}), true},
		{SliceToList([]int{2}), true},
		{SliceToList([]int{1, 2, 1, 4, 1, 2, 1}), true},
		{SliceToList([]int{1, 2, 1, 4, 4, 1, 2, 1}), true},
		{SliceToList([]int{1, 2, 1, 4, 5, 1, 2, 1}), false},
		{SliceToList([]int{1, 2, 5, 5, 1, 2, 1}), false},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := IsPalindrome(c.input)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v",
					c.expect, got, ListToSlice(c.input))
			}
		})
	}
}
