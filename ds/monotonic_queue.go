package ds

// MonotonicQueue 单调队列 队列中的数据单调递增 主要用来辅助解决滑动窗口相关的问题
type MonotonicQueue struct {
	queue []int
}

func NewMonotonicQueue() *MonotonicQueue {
	return &MonotonicQueue{nil}
}

func (s *MonotonicQueue) Size() int {
	return len(s.queue)
}

func (s *MonotonicQueue) Peek() int {
	return s.queue[0]
}

func (s *MonotonicQueue) Dequeue(head int) bool {
	res := s.Size() != 0 && s.queue[0] == head
	if res {
		s.queue = s.queue[1:]
	}
	return res
}

func (s *MonotonicQueue) Enqueue(value int) {
	for s.Size() != 0 && value > s.queue[s.Size()-1] {
		s.queue = s.queue[:s.Size()-1]
	}
	s.queue = append(s.queue, value)
}
