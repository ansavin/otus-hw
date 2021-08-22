package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("simple push front", func(t *testing.T) {
		l := NewList()

		l.PushFront(10)

		require.Equal(t, 1, l.Len())
		front := l.Front()
		back := l.Back()
		require.Equal(t, 10, front.Value)
		require.Equal(t, 10, back.Value)
		require.Nil(t, front.Prev)
		require.Nil(t, front.Next)
		require.Nil(t, back.Prev)
		require.Nil(t, back.Next)
	})

	t.Run("simple push back", func(t *testing.T) {
		l := NewList()

		l.PushBack(10)

		require.Equal(t, 1, l.Len())
		front := l.Front()
		back := l.Back()
		require.Equal(t, 10, front.Value)
		require.Equal(t, 10, back.Value)
		require.Nil(t, front.Prev)
		require.Nil(t, front.Next)
		require.Nil(t, back.Prev)
		require.Nil(t, back.Next)
	})

	t.Run("simple remove", func(t *testing.T) {
		l := NewList()

		l.PushBack(10)
		l.PushBack(20)
		l.PushBack(30)

		middle := l.Front().Next
		l.Remove(middle)

		require.Equal(t, 2, l.Len())
		front := l.Front()
		back := l.Back()
		require.Equal(t, 10, front.Value)
		require.Equal(t, 30, front.Next.Value)
		require.Equal(t, 30, back.Value)
		require.Equal(t, 10, back.Prev.Value)

		l.Remove(back)
		require.Equal(t, 1, l.Len())
		front = l.Front()
		require.Equal(t, 10, front.Value)
		require.Nil(t, front.Prev)
		require.Nil(t, front.Next)

		l.Remove(front)
		require.Equal(t, 0, l.Len())
		front = l.Front()
		require.Nil(t, front)
	})

	t.Run("simple remove", func(t *testing.T) {
		l := NewList()

		l.PushBack(10) // [10]
		l.PushBack(20) // [10,20]
		l.PushBack(30) // [10,20,30]

		middle := l.Front().Next
		l.MoveToFront(middle) // [20,10,30]

		require.Equal(t, 3, l.Len())
		front := l.Front()
		back := l.Back()
		require.Equal(t, 20, front.Value)
		require.Equal(t, 10, front.Next.Value)
		require.Equal(t, 30, back.Value)
		require.Equal(t, 10, back.Prev.Value)

		l.MoveToFront(back) // [30,20,10]
		require.Equal(t, 3, l.Len())
		front = l.Front()
		back = l.Back()
		require.Equal(t, 30, front.Value)
		require.Equal(t, 20, front.Next.Value)
		require.Equal(t, 10, back.Value)
		require.Equal(t, 20, back.Prev.Value)
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})
}
