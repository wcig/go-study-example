package simplequeue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	q := New()
	assert.Equal(t, 0, q.Size())
	assert.True(t, q.IsEmpty())

	v, ok := q.Peek()
	assert.Nil(t, v)
	assert.False(t, ok)

	for i := 0; i < 5; i++ {
		q.Push(i)
	}
	assert.Equal(t, 5, q.Size())
	assert.False(t, q.IsEmpty())
	assert.Equal(t, []interface{}{0, 1, 2, 3, 4}, q.Values())

	for i := 0; i < 5; i++ {
		v, ok = q.Peek()
		assert.Equal(t, i, v)
		assert.True(t, ok)
		v, ok = q.Pop()
		assert.Equal(t, i, v)
		assert.True(t, ok)
	}

	v, ok = q.Peek()
	assert.Nil(t, v)
	assert.False(t, ok)

	assert.Equal(t, 0, q.Size())
	assert.True(t, q.IsEmpty())
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
		q.Push(i)
	}
	// 注意一般这时候需重新获取新创建迭代器
	iterator = q.Iterator()
	for iterator.HasNext() {
		values = append(values, iterator.Next())
	}
	assert.Equal(t, q.Values(), values)
	assert.Equal(t, []interface{}{0, 1, 2, 3, 4}, values)
}
