package convert_binary_number_in_a_linked_list_to_integer

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestGetDecimalValue(t *testing.T) {
	cases := []struct {
		input  *ListNode
		expect int
	}{
		{SliceToList([]int{1, 0, 1}), 5},
		{SliceToList([]int{0}), 0},
		{SliceToList([]int{1}), 1},
		{SliceToList([]int{1, 0, 0, 1, 0, 0, 1, 1, 1, 0, 0, 0, 0, 0, 0}), 18880},
		{SliceToList([]int{0, 0}), 0},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := GetDecimalValue(c.input)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v",
					c.expect, got, ListToSlice(c.input))
			}
		})
	}
}
