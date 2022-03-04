package arraystack

import (
	"testing"

	"github.com/emirpasic/gods/stacks/arraystack"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	s := New()
	assert.Equal(t, 0, s.Size())
	assert.True(t, s.IsEmpty())
	v, ok := s.Peek()
	assert.Nil(t, v)
	assert.False(t, ok)

	for i := 0; i < 5; i++ {
		s.Push(i)
	}
	assert.Equal(t, 5, s.Size())
	assert.False(t, s.IsEmpty())
	assert.Equal(t, []interface{}{4, 3, 2, 1, 0}, s.Values())

	for i := 4; i >= 0; i-- {
		v, ok = s.Peek()
		assert.Equal(t, i, v)
		assert.True(t, ok)
		v, ok = s.Pop()
		assert.Equal(t, i, v)
		assert.True(t, ok)
	}

	v, ok = s.Peek()
	assert.Nil(t, v)
	assert.False(t, ok)
	assert.Equal(t, 0, s.Size())
	assert.True(t, s.IsEmpty())
	assert.Equal(t, []interface{}{}, s.Values())
}

func TestIterator(t *testing.T) {
	s := New()
	iterator := s.Iterator()

	var values []interface{}
	for iterator.HasNext() {
		values = append(values, iterator.Next())
	}
	assert.True(t, len(values) == 0)

	for i := 0; i < 5; i++ {
		s.Push(i)
	}
	for iterator.HasNext() {
		values = append(values, iterator.Next())
	}
	assert.Equal(t, s.Values(), values)
	assert.Equal(t, []interface{}{4, 3, 2, 1, 0}, values)
}

func TestOrder(t *testing.T) {
	s1 := New()
	s2 := arraystack.New()
	for i := 0; i < 5; i++ {
		s1.Push(i)
		s2.Push(i)
	}
	assert.Equal(t, s2.Values(), s1.Values())

	var v1 []interface{}
	var v2 []interface{}
	i1 := s1.Iterator()
	for i1.HasNext() {
		v1 = append(v1, i1.Next())
	}
	i2 := s2.Iterator()
	for i2.Next() {
		v2 = append(v2, i2.Value())
	}
	assert.Equal(t, v2, v1)
}
