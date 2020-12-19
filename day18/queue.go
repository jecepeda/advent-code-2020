package day18

type Queue struct {
	Data []Token
}

func (q *Queue) Dequeue() Token {
	if len(q.Data) == 0 {
		return Token{}
	}
	v := q.Data[0]
	q.Data = q.Data[1:]
	return v
}

func (q *Queue) Peek() Token {
	if len(q.Data) == 0 {
		return Token{}
	}
	return q.Data[0]
}

func (q *Queue) Len() int {
	return len(q.Data)
}

func (q *Queue) Add(t Token) {
	q.Data = append(q.Data, t)
}
