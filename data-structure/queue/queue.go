package queue

type Queue interface {
	Push(d int)
	Pop() int
	GetSize() int
	IsEmpty() bool
}
