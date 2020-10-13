package libs

type queueItem interface{}
type Queue []queueItem
type QueueInterface interface {
	Push(item queueItem)
	Pop()
	IsEmpty()
	Length()
	Top()
}

func (q *Queue) Push(item queueItem) {
	*q = append(*q, item)
}
func (q *Queue) Top() queueItem {
	return (*q)[0]
}
func (q *Queue) Pop() queueItem {
	top := q.Top()
	*q = (*q)[1:]
	return top
}
func (q *Queue) Length() int {
	return len(*q)
}
func (q Queue) IsEmpty() bool {
	return q.Length() == 0
}
