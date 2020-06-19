package IntersectionOfTwoLists

import (
	. "github.com/bumblebee211996/go-ds-daa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

type testcase struct {
	input1 *ListNode
	input2 *ListNode
	expect int
}

func MakeIntersectionList(head1, head2 *ListNode) (*ListNode, *ListNode) {
	c2 := head2
	for ; c2.Next != nil; c2 = c2.Next {
	}
	slow, fast := head1, head1
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	c2.Next = slow
	return head1, head2
}

func TestIntersectingNode1(t *testing.T) {
	cases := make([]testcase, 0)
	h1, h2 := MakeIntersectionList(SliceToList([]int{1, 2, 3, 4, 5, 6}), SliceToList([]int{7, 8, 9}))
	cases = append(cases, testcase{input1: h1, input2: h2, expect: 4})
	h1, h2 = MakeIntersectionList(SliceToList([]int{1, 2, 3, 4, 5}), SliceToList([]int{7, 9}))
	cases = append(cases, testcase{input1: h1, input2: h2, expect: 3})
	h1, h2 = MakeIntersectionList(SliceToList([]int{3, 4, 5, 6}), SliceToList([]int{7, 8}))
	cases = append(cases, testcase{input1: h1, input2: h2, expect: 5})
	h1, h2 = SliceToList([]int{3, 4, 5, 6}), SliceToList([]int{7, 8})
	cases = append(cases, testcase{input1: h1, input2: h2, expect: -1})
	h1, h2 = SliceToList([]int{3, 4, 5, 6}), SliceToList([]int{})
	cases = append(cases, testcase{input1: h1, input2: h2, expect: -1})
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := IntersectingNode1(c.input1, c.input2)
			if got != nil && !reflect.DeepEqual(got.Val, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v %v",
					c.expect, got, c.input1, c.input2)
			}
		})
	}
}

func TestIntersectingNode2(t *testing.T) {
	cases := make([]testcase, 0)
	h1, h2 := MakeIntersectionList(SliceToList([]int{1, 2, 3, 4, 5, 6}), SliceToList([]int{7, 8, 9}))
	cases = append(cases, testcase{input1: h1, input2: h2, expect: 4})
	h1, h2 = MakeIntersectionList(SliceToList([]int{1, 2, 3, 4, 5}), SliceToList([]int{7, 9}))
	cases = append(cases, testcase{input1: h1, input2: h2, expect: 3})
	h1, h2 = MakeIntersectionList(SliceToList([]int{3, 4, 5, 6}), SliceToList([]int{7, 8}))
	cases = append(cases, testcase{input1: h1, input2: h2, expect: 5})
	h1, h2 = SliceToList([]int{3, 4, 5, 6}), SliceToList([]int{7, 8})
	cases = append(cases, testcase{input1: h1, input2: h2, expect: -1})
	h1, h2 = SliceToList([]int{3, 4, 5, 6}), SliceToList([]int{})
	cases = append(cases, testcase{input1: h1, input2: h2, expect: -1})
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := IntersectingNode2(c.input1, c.input2)
			if got != nil && !reflect.DeepEqual(got.Val, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v %v",
					c.expect, got, c.input1, c.input2)
			}
		})
	}
}

func TestIntersectingNode3(t *testing.T) {
	cases := make([]testcase, 0)
	h1, h2 := MakeIntersectionList(SliceToList([]int{1, 2, 3, 4, 5, 6}), SliceToList([]int{7, 8, 9}))
	cases = append(cases, testcase{input1: h1, input2: h2, expect: 4})
	h1, h2 = MakeIntersectionList(SliceToList([]int{1, 2, 3, 4, 5}), SliceToList([]int{7, 9}))
	cases = append(cases, testcase{input1: h1, input2: h2, expect: 3})
	h1, h2 = MakeIntersectionList(SliceToList([]int{3, 4, 5, 6}), SliceToList([]int{7, 8}))
	cases = append(cases, testcase{input1: h1, input2: h2, expect: 5})
	h1, h2 = SliceToList([]int{3, 4, 5, 6}), SliceToList([]int{7, 8})
	cases = append(cases, testcase{input1: h1, input2: h2, expect: -1})
	h1, h2 = SliceToList([]int{3, 4, 5, 6}), SliceToList([]int{})
	cases = append(cases, testcase{input1: h1, input2: h2, expect: -1})
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := IntersectingNode3(c.input1, c.input2)
			if got != nil && !reflect.DeepEqual(got.Val, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v %v",
					c.expect, got, c.input1, c.input2)
			}
		})
	}
}

func TestIntersectingNode4(t *testing.T) {
	cases := make([]testcase, 0)
	h1, h2 := MakeIntersectionList(SliceToList([]int{1, 2, 3, 4, 5, 6}), SliceToList([]int{7, 8, 9}))
	cases = append(cases, testcase{input1: h1, input2: h2, expect: 4})
	h1, h2 = MakeIntersectionList(SliceToList([]int{1, 2, 3, 4, 5}), SliceToList([]int{7, 9}))
	cases = append(cases, testcase{input1: h1, input2: h2, expect: 3})
	h1, h2 = MakeIntersectionList(SliceToList([]int{3, 4, 5, 6}), SliceToList([]int{7, 8}))
	cases = append(cases, testcase{input1: h1, input2: h2, expect: 5})
	h1, h2 = SliceToList([]int{3, 4, 5, 6}), SliceToList([]int{7, 8})
	cases = append(cases, testcase{input1: h1, input2: h2, expect: -1})
	h1, h2 = SliceToList([]int{3, 4, 5, 6}), SliceToList([]int{})
	cases = append(cases, testcase{input1: h1, input2: h2, expect: -1})
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := IntersectingNode4(c.input1, c.input2)
			if got != nil && !reflect.DeepEqual(got.Val, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v %v",
					c.expect, got, c.input1, c.input2)
			}
		})
	}
}
