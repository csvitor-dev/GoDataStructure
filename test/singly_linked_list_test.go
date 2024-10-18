package test

import (
	. "GoDataStructure/collections"
	"testing"
)

func Test_VerifyEmptyList(t *testing.T) {
	ll := NewSinglyLinkedList[int]()
	_, err := ll.Delete()

	if err == nil {
		ErrorMessage(t, "deleting elements does not behave well", nil, err)
	}
	_, err = ll.RemoveAt(3)

	if err == nil {
		ErrorMessage(t, "deleting elements does not behave well", nil, err)
	}
}

func Test_FindElements(t *testing.T) {
	cases := []struct {
		got Input
		expected Output
	}{
		{
			Input{ []int{ 2, 7, 5 }, 0, },
			Output{ 2, false, },
		},
		{
			Input{ []int{ 2, 10, 8, 19 }, 2, },
			Output{ 8, false, },
		},
		{
			Input{ []int{}, 9, },
			Output{ 0, true, },
		},
	}

	for _, tc := range cases {
		ll := NewSinglyLinkedList[int]()

		for _, v := range tc.got.Values {
			ll.Add(v)
		}

		if v, err := FindOnIndex(ll, tc.got.Index); 
			(err != nil) != tc.expected.Err || v != tc.expected.Value {
			ErrorMessage(t, "search for element on index failed", tc.expected.Value, v)
		}
	}
}

func Test_AddTwoElements(t *testing.T) {
	ll := NewSinglyLinkedList[int]()
	ll.Add(3)
	ll.Add(5)
	
	if length := ll.Length(); length != 2 {
		ErrorMessage(t, "adding elements misbehaves", 2, length)
	}
}

func Test_InsertElementsByIndex(t *testing.T) {

}

func Test_DeleteElements(t *testing.T) {
	
}

func Test_RemoveElementsByIndex(t *testing.T) {
	
}