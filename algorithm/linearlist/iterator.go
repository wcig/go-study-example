package linearlist

type Iterator interface {
	HasNext() bool
	Next() interface{}
}
