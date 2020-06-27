package flatten_multilevel_doubly_linked_list

import (
	"reflect"
	"strconv"
	"testing"
)

type testcases struct {
	input  *Node
	expect *Node
}

func sliceToList(values []int) *Node {
	head := &Node{Val: -1}
	temp := head
	for _, val := range values {
		temp.Next = &Node{Val: val, Prev: temp}
		temp = temp.Next
	}
	return head.Next
}

func listToSlice(head *Node) []int {
	values := make([]int, 0)
	for temp := head; temp != nil; temp = temp.Next {
		values = append(values, temp.Val)
	}
	return values
}

func TestFlatten(t *testing.T) {
	cases := make([]testcases, 0)
	l1, l2, l3 := sliceToList([]int{1, 2, 3, 4, 5, 6}), sliceToList([]int{7, 8, 9, 10}), sliceToList([]int{11, 12})
	l2.Next.Child = l3
	l1.Next.Next.Child = l2
	cases = append(cases, testcases{l1, sliceToList([]int{1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6})})
	l1, l2 = sliceToList([]int{1, 2}), sliceToList([]int{3})
	l1.Child = l2
	cases = append(cases, testcases{l1, sliceToList([]int{1, 3, 2})})
	cases = append(cases, testcases{nil, nil})
	l1, l2 = sliceToList([]int{1, 2}), sliceToList([]int{3})
	l1.Next.Child = l2
	cases = append(cases, testcases{l1, sliceToList([]int{1, 2, 3})})
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Flatten(c.input)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v",
					listToSlice(c.expect), listToSlice(got), listToSlice(c.input))
			}
		})
	}
}
