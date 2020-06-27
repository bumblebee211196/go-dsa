package partition_list

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"reflect"
	"strconv"
	"testing"
)

func TestPartition(t *testing.T) {
	cases := []struct {
		input1 *ListNode
		input2 int
		expect *ListNode
	}{
		{SliceToList([]int{}), 3, SliceToList([]int{})},
		{SliceToList([]int{1, 4, 3, 2, 5, 2}), 3, SliceToList([]int{1, 2, 2, 4, 3, 5})},
		{SliceToList([]int{1, 2, 3, 4, 5, 6}), 10, SliceToList([]int{1, 2, 3, 4, 5, 6})},
		{SliceToList([]int{8, 9, 7, 4, 2, 6, 3, 1}), 7, SliceToList([]int{4, 2, 6, 3, 1, 8, 9, 7})},
	}
	for i, c := range cases {
		t.Run("TestCase"+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Partition(c.input1, c.input2)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("Expected: %v, got %v for input %v %v",
					ListToSlice(c.expect), ListToSlice(got), ListToSlice(c.input1), c.input2)
			}
		})
	}
}
