package simplequeue

import (
	"fmt"
	"sync"
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

func TestConcurrent(t *testing.T) {
	q := New()

	var wg sync.WaitGroup
	wg.Add(5)
	var expectSum int
	for i := 0; i < 5; i++ {
		v := i
		expectSum += v
		go func() {
			q.Push(v)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(q.Values())
	assert.Equal(t, 5, q.size)
	var n int
	for _, v := range q.Values() {
		n += v.(int)
	}
	fmt.Println(n)
	assert.Equal(t, expectSum, n)
}
