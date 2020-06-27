package delete_node_in_linked_list

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

type testcase struct {
	list   *ListNode
	input  *ListNode
	expect *ListNode
}

func TestDeleteNode(t *testing.T) {
	cases := make([]testcase, 0)
	h, e := SliceToList([]int{1, 2, 3, 3, 4}), SliceToList([]int{2, 3, 3, 4})
	cases = append(cases, testcase{h, h, e})
	h, e = SliceToList([]int{1, 2, 3, 3, 4}), SliceToList([]int{1, 3, 3, 4})
	cases = append(cases, testcase{h, h.Next, e})
	h, e = SliceToList([]int{1, 2, 3, 3, 4}), SliceToList([]int{1, 2, 3, 4})
	cases = append(cases, testcase{h, h.Next.Next, e})
	h, e = SliceToList([]int{1, 2, 3, 3, 4}), SliceToList([]int{1, 2, 3, 4})
	cases = append(cases, testcase{h, h.Next.Next.Next, e})
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			DeleteNode(c.input)
			if !reflect.DeepEqual(c.list, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v",
					ListToSlice(c.expect), ListToSlice(c.list), ListToSlice(c.input))
			}
		})
	}
}
