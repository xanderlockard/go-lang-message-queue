package queue

type Queable struct {
	queueName     string
	processId     int
	processStatus string
	next          *Queable
}

type LinkedListQueue struct {
	first *Queable
	last  *Queable
	size  int
}

type Queue interface {
	Enqueue(q Queable) error
	Dequeue() (Queable, error)
	Peek() (Queable, error)
	Size() int
	Last() (Queable, error)
}

func (q *LinkedListQueue) Enqueue(qb Queable) error {
	current := q.Last()
	if current == nil || current.processId == 0 {
		q.first = &qb
		q.last = &qb
		q.size++
		return nil
	} else if current.processId != 0 {
		current.next = &qb
		q.last = &qb
	}
	q.size++
	return nil
}

func (q *LinkedListQueue) Dequeue() (*Queable, error) {
	current := q.First()
	if current == nil || current.processId == 0 {
		return nil, nil
	}
	if current.next != nil {
		q.first = current.next
		q.size--
		return current, nil
	} else {
		q.first = nil
		q.last = nil
		q.size--
		return current, nil
	}
}

func (q *LinkedListQueue) Size() int {
	return q.size
}

func (q *LinkedListQueue) First() *Queable {
	if q.Size() == 0 {
		return nil
	}
	return q.first
}

func (q *LinkedListQueue) Last() *Queable {
	if q.Size() == 0 {
		return nil
	}
	return q.last
}
