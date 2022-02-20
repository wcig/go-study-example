package singlelinklist

// 单链表
type SingleLinkList struct {
	first *node
	size  int
}

type node struct {
	data interface{}
	next *node
}

func New() *SingleLinkList {
	return &SingleLinkList{
		first: nil,
		size:  0,
	}
}
