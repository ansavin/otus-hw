package hw04lrucache

import "fmt"

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type Scroll struct {
	len   int
	first *ListItem
	last  *ListItem
}

func (s *Scroll) Len() int {
	return s.len
}

func (s *Scroll) Front() *ListItem {
	return s.first
}

func (s *Scroll) Back() *ListItem {
	return s.last
}

func (s *Scroll) PushBack(v interface{}) *ListItem {
	d := ListItem{
		Value: v,
		Prev:  s.last,
		Next:  nil,
	}
	if s.len == 0 {
		s.first = &d
	}
	if s.last != nil {
		s.last.Next = &d
	}
	s.len++
	s.last = &d
	return &d
}

func (s *Scroll) PushFront(v interface{}) *ListItem {
	d := ListItem{
		Value: v,
		Prev:  nil,
		Next:  s.first,
	}
	if s.len == 0 {
		s.last = &d
	}
	if s.first != nil {
		s.first.Prev = &d
	}
	s.len++
	s.first = &d
	return &d
}

func (s *Scroll) Remove(i *ListItem) {
	if i == nil {
		fmt.Println("cant delete nil")
		return
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		if i.Next != nil {
			i.Next.Prev = nil
		}
		s.first = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		if i.Prev != nil {
			i.Prev.Next = nil
		}
		s.last = i.Prev
	}
	s.len--
}

func (s *Scroll) MoveToFront(i *ListItem) {
	if i == nil {
		return
	}
	s.Remove(i)
	s.PushFront(i.Value)
}

func NewList() *Scroll {
	return new(Scroll)
}
