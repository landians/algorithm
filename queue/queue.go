package queue

// Queue 队列的实现
type Queue struct {
	Values []interface{}
}

func New(size int) *Queue {
	return &Queue{
		Values: make([]interface{}, 0, size),
	}
}

// Enqueue 入队
func (q *Queue) Enqueue(v interface{}) {
	if q.IsFull() {
		return
	}
	q.Values = append(q.Values, v)
}

// Dequeue 出队
func (q *Queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	v := q.Values[0]
	q.Values = q.Values[1:]
	return v
}

// IsFull 判断队列满
func (q *Queue) IsFull() bool {
	return cap(q.Values) == len(q.Values)
}

// IsEmpty 判断队列空
func (q *Queue) IsEmpty() bool {
	return len(q.Values) == 0
}
