package add_two_numbers

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestAddTwoNumbers1(t *testing.T) {
	cases := []struct {
		input1 *ListNode
		input2 *ListNode
		expect *ListNode
	}{
		{SliceToList([]int{7, 2, 4, 3}), SliceToList([]int{5, 6, 4}), SliceToList([]int{7, 8, 0, 7})},
		{SliceToList([]int{9, 9, 9}), SliceToList([]int{1}), SliceToList([]int{1, 0, 0, 0})},
		{SliceToList([]int{9}), SliceToList([]int{9, 9, 1}), SliceToList([]int{1, 0, 0, 0})},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := AddTwoNumbers1(c.input1, c.input2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v %v",
					ListToSlice(c.expect), ListToSlice(got), ListToSlice(c.input1), ListToSlice(c.input2))
			}
		})
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	cases := []struct {
		input1 *ListNode
		input2 *ListNode
		expect *ListNode
	}{
		{SliceToList([]int{7, 2, 4, 3}), SliceToList([]int{5, 6, 4}), SliceToList([]int{7, 8, 0, 7})},
		{SliceToList([]int{9, 9, 9}), SliceToList([]int{1}), SliceToList([]int{1, 0, 0, 0})},
		{SliceToList([]int{9}), SliceToList([]int{9, 9, 1}), SliceToList([]int{1, 0, 0, 0})},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := AddTwoNumbers2(c.input1, c.input2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v %v",
					c.expect, got, c.input1, c.input2)
			}
		})
	}
}
