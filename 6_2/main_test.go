package main

import (
	"testing"
)

func Test(t *testing.T) {
	var (
		list  IntList
		item1 IntItem
		item2 IntItem
		item3 IntItem
	)

	// Initial
	list = IntList{}
	if len := list.Len(); len != 0 {
		t.Fatalf("Empty list should be with len=0, but it is %v\n", len)
	}

	if item := list.First(); item != nil {
		t.Fatalf("Empty list should contain first=nil, but it is %v\n", item)
	}

	if item := list.Last(); item != nil {
		t.Fatalf("Empty list should contain last=nil, but it is %v\n", item)
	}

	// First item
	list = IntList{}
	list.PushFront(&IntItem{})

	if len := list.Len(); len != 1 {
		t.Fatalf("List with single item should be len=1, but it is %v\n", len)
	}

	if first, last := list.First(), list.Last(); first != last {
		t.Fatalf("List with single item should contain equal first and last items, but they are %v and %v\n", first, last)
	}

	list = IntList{}
	list.PushBack(&IntItem{})

	if len := list.Len(); len != 1 {
		t.Fatalf("List with single item should be len=1, but it is %v\n", len)
	}

	if first, last := list.First(), list.Last(); first != last {
		t.Fatalf("List with single item should contain equal first and last items, but they are %v and %v\n", first, last)
	}

	// Some items
	list = IntList{}
	list.PushFront(&IntItem{value: 2})
	list.PushFront(&IntItem{value: 1})
	list.PushBack(&IntItem{value: 3})

	if len := list.Len(); len != 3 {
		t.Fatalf("List should be len=3, but it is %v\n", len)
	}

	if first := list.First(); first.Value() != 1 {
		t.Fatalf("First item should contain value equal 1, but it is %v\n", first.Value())
	}

	if second := list.First().next; second.Value() != 2 {
		t.Fatalf("Second item should contain value equal 2, but it is %v\n", second.Value())
	}

	if last := list.Last(); last.Value() != 3 {
		t.Fatalf("Last item should contain value equal 3, but it is %v\n", last.Value())
	}

	// Removing of first item
	list = IntList{}
	item1 = IntItem{value: 1}
	item2 = IntItem{value: 2}
	item3 = IntItem{value: 3}
	list.PushBack(&item1)
	list.PushBack(&item2)
	list.PushBack(&item3)

	list.Remove(&item1)

	if len := list.Len(); len != 2 {
		t.Fatalf("List should be len=2, but it is %v\n", len)
	}

	if first := list.First(); first.Value() != item2.Value() {
		t.Fatalf("First item should contain value equal %d, but it is %v\n", first.Value(), item2.Value())
	}

	if second := list.First().next; second.Value() != item3.Value() {
		t.Fatalf("Second item should contain value equal %d, but it is %v\n", second.Value(), item3.Value())
	}

	// Removing of last item
	list = IntList{}
	item1 = IntItem{value: 1}
	item2 = IntItem{value: 2}
	item3 = IntItem{value: 3}
	list.PushBack(&item1)
	list.PushBack(&item2)
	list.PushBack(&item3)

	list.Remove(&item3)

	if len := list.Len(); len != 2 {
		t.Fatalf("List should be len=2, but it is %v\n", len)
	}

	if second := list.Last(); second.Value() != item2.Value() {
		t.Fatalf("Second item should contain value equal %d, but it is %v\n", second.Value(), item2.Value())
	}

	if first := list.First(); first.Value() != item1.Value() {
		t.Fatalf("First item should contain value equal %d, but it is %v\n", first.Value(), item1.Value())
	}

	// Removing of middle item
	list = IntList{}
	item1 = IntItem{value: 1}
	item2 = IntItem{value: 2}
	item3 = IntItem{value: 3}
	list.PushBack(&item1)
	list.PushBack(&item2)
	list.PushBack(&item3)

	list.Remove(&item2)

	if len := list.Len(); len != 2 {
		t.Fatalf("List should be len=2, but it is %v\n", len)
	}

	if first := list.First(); first.Value() != item1.Value() {
		t.Fatalf("First item should contain value equal %d, but it is %v\n", first.Value(), item2.Value())
	}

	if second := list.First().next; second.Value() != item3.Value() {
		t.Fatalf("Second item should contain value equal %d, but it is %v\n", second.Value(), item3.Value())
	}
}
