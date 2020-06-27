package remove_linked_list_elements

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	cases := []struct {
		list   *ListNode
		val    int
		expect *ListNode
	}{
		{SliceToList([]int{1, 2, 6, 3, 4, 5, 6}), 6, SliceToList([]int{1, 2, 3, 4, 5})},
		{SliceToList([]int{1, 2, 2, 2, 3}), 2, SliceToList([]int{1, 3})},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := RemoveElements(CopyList(c.list), c.val)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v %v",
					ListToSlice(c.expect), ListToSlice(got), ListToSlice(c.list), c.val)
			}
		})
	}
}
