package single_call

import (
	"sync/atomic"
)

type SingleCall struct {
	doing int32
}

func (s *SingleCall) Do(f func()) {
	if atomic.CompareAndSwapInt32(&s.doing, 0, 1) {
		f()
		_ = atomic.CompareAndSwapInt32(&s.doing, 1, 0)
	}
}
