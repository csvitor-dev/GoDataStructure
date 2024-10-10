package test

import (
	. "GoDataStructure/collections"
	"testing"
)

func Test_EmptyList(t *testing.T) {
	ll := NewSinglyLinkedList[int]()
	_, err := ll.Delete()

	if err == nil {
		t.Errorf(message("Deleting elements does not behave well", nil, err))
	}
	_, err = ll.RemoveAt(3)

	if err == nil {
		t.Error(message("Deleting elements does not behave well", nil, err))
	}
}

func Test_FindElements(t *testing.T) {

}

func Test_AddElements(t *testing.T) {
	
	ll := NewSinglyLinkedList[int]()
	ll.Add(0)
	ll.Add(5)
	Print(ll)

	if ll.Length() != 2 {
		t.Error(message("Adding elements misbehaves", 2, ll.Length()))
	}

	if v, err := FindOnIndex(ll, 2); err != nil || v != 5 {
		t.Error(message("Constant insertion fail", 5, v))
	}
}

func Test_InsertElementsByIndex(t *testing.T) {

}

func Test_DeleteElements(t *testing.T) {
	
}

func Test_RemoveElementsByIndex(t *testing.T) {
	
}