package MiddleOfLinkedList

import (
	. "github.com/bumblebee211996/go-ds-daa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestMiddleOfLinkedList1(t *testing.T) {
	cases := []struct {
		input  *ListNode
		expect int
	}{
		{SliceToList([]int{}), -999},
		{SliceToList([]int{1}), 1},
		{SliceToList([]int{1, 2, 3}), 2},
		{SliceToList([]int{1, 2, 3, 4}), 3},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := MidOfLinkedList1(c.input)
			if got != nil && c.expect != -999 && !reflect.DeepEqual(got.Val, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v",
					c.expect, got, c.input)
			}
		})
	}
}

func TestMiddleOfLinkedList2(t *testing.T) {
	cases := []struct {
		input  *ListNode
		expect int
	}{
		{SliceToList([]int{}), -999},
		{SliceToList([]int{1}), 1},
		{SliceToList([]int{1, 2, 3}), 2},
		{SliceToList([]int{1, 2, 3, 4}), 3},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := MidOfLinkedList2(c.input)
			if got != nil && c.expect != -999 && !reflect.DeepEqual(got.Val, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v",
					c.expect, got, c.input)
			}
		})
	}
}
