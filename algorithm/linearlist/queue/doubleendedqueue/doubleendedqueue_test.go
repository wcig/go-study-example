package doubleendedqueue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushBack(t *testing.T) {
	q := New()
	for i := 0; i < 5; i++ {
		q.PushBack(i)
	}
	values := []interface{}{0, 1, 2, 3, 4}
	assert.Equal(t, len(values), q.size)
	assert.Equal(t, values, q.Values())
}

func TestPushFront(t *testing.T) {
	q := New()
	for i := 0; i < 5; i++ {
		q.PushFront(i)
	}
	values := []interface{}{4, 3, 2, 1, 0}
	assert.Equal(t, len(values), q.size)
	assert.Equal(t, values, q.Values())
}

func TestPopBack(t *testing.T) {
	q := New()
	value, ok := q.PopBack()
	assert.Nil(t, value)
	assert.False(t, ok)

	var exceptValues []interface{}
	for i := 0; i < 5; i++ {
		q.PushFront(i)
		exceptValues = append(exceptValues, i)
	}
	fmt.Println(q.Values()) // [4 3 2 1 0]

	var values []interface{}
	for i := 0; i < 5; i++ {
		value, ok = q.PopBack()
		assert.NotNil(t, value)
		assert.True(t, ok)
		values = append(values, value)
	}
	assert.Equal(t, exceptValues, values)

	value, ok = q.PopBack()
	assert.Nil(t, value)
	assert.False(t, ok)
}

func TestPopFront(t *testing.T) {
	q := New()
	value, ok := q.PopFront()
	assert.Nil(t, value)
	assert.False(t, ok)

	var exceptValues []interface{}
	for i := 0; i < 5; i++ {
		q.PushBack(i)
		exceptValues = append(exceptValues, i)
	}
	fmt.Println(q.Values()) // [0 1 2 3 4]

	var values []interface{}
	for i := 0; i < 5; i++ {
		value, ok = q.PopFront()
		assert.NotNil(t, value)
		assert.True(t, ok)
		values = append(values, value)
	}
	assert.Equal(t, exceptValues, values)

	value, ok = q.PopFront()
	assert.Nil(t, value)
	assert.False(t, ok)
}

func TestSimple(t *testing.T) {
	q := New()
	assert.True(t, q.IsEmpty())
	assert.Equal(t, 0, q.size)
	assert.Equal(t, []interface{}{}, q.Values())

	for i := 0; i < 5; i++ {
		q.PushBack(i)
	}
	assert.False(t, q.IsEmpty())
	assert.Equal(t, 5, q.size)
	assert.Equal(t, []interface{}{0, 1, 2, 3, 4}, q.Values())

	q.Clear()
	assert.True(t, q.IsEmpty())
	assert.Equal(t, 0, q.size)
	assert.Equal(t, []interface{}{}, q.Values())
}

func TestIterator(t *testing.T) {
	q := New()
	iterator := q.Iterator()

	var values []interface{}
	for iterator.HasNext() {
		values = append(values, iterator.Next())
	}
	assert.True(t, len(values) == 0)

	for i := 0; i < 5; i++ {
		q.PushBack(i)
	}
	iterator = q.Iterator()
	for iterator.HasNext() {
		values = append(values, iterator.Next())
	}
	assert.Equal(t, q.Values(), values)
	assert.Equal(t, []interface{}{0, 1, 2, 3, 4}, values)

	q.Clear()
	values = []interface{}{}
	for i := 0; i < 5; i++ {
		q.PushFront(i)
	}
	iterator = q.Iterator()
	for iterator.HasNext() {
		values = append(values, iterator.Next())
	}
	assert.Equal(t, q.Values(), values)
	assert.Equal(t, []interface{}{4, 3, 2, 1, 0}, values)
}
