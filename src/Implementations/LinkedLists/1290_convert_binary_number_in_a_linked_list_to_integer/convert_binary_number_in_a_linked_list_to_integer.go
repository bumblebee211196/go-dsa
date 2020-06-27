package convert_binary_number_in_a_linked_list_to_integer

import (
	. "github.com/bumblebee211996/go-dsa/src/Implementations/LinkedLists"
	"strconv"
	"strings"
)

// Problem Statement: Given head which is a reference node to a singly-linked list.
// The value of each node in the linked list is either 0 or 1. The linked list holds
// the binary representation of a number. Return the decimal value of the number in
// the linked list.

// Solution: Straight forward approach
// We build a string from the list. The order of the integers in the string will be
// from head to tail. We then use the inbuilt "strconv" package to convert the binary
// string to integer
func GetDecimalValue(head *ListNode) int {
	var bin strings.Builder
	for head != nil {
		bin.WriteString(strconv.Itoa(head.Val))
		head = head.Next
	}
	i, _ := strconv.ParseInt(bin.String(), 2, 64)
	return int(i)
}
