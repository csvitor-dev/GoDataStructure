package test

import (
	. "GoDataStructure/collections"
	"testing"
)

func Test_EmptyList(t *testing.T) {
	ll := NewSinglyLinkedList[int]()
	_, err := ll.Delete()

	if err == nil {
		t.Errorf(message("deleting elements does not behave well", nil, err))
	}
	_, err = ll.RemoveAt(3)

	if err == nil {
		t.Error(message("deleting elements does not behave well", nil, err))
	}
}

func Test_FindElements(t *testing.T) {

}

func Test_AddElements(t *testing.T) {
	
	ll := NewSinglyLinkedList[int]()
	ll.Add(3)
	ll.Add(5)

	if ll.Length() != 2 {
		t.Error(message("adding elements misbehaves", 2, ll.Length()))
	}

	if v, err := FindOnIndex(ll, 1); err != nil || v != 5 {
		t.Error(message("constant insertion fail", 5, v))
	}
}

func Test_InsertElementsByIndex(t *testing.T) {

}

func Test_DeleteElements(t *testing.T) {
	
}

func Test_RemoveElementsByIndex(t *testing.T) {
	
}