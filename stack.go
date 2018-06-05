package main

type Stack struct {
	in, out chan []byte
	pop     chan struct{}
	len     int

	data [][]byte
}

func NewStack() *Stack {
	stack := &Stack{
		in:  make(chan []byte),
		out: make(chan []byte),
		pop: make(chan struct{}),
		// len(s.data) is eventually consistent due to goroutine scheduling
		// s.len is always consistent
		len: 0,
	}
	go stack.run()
	return stack
}

func (s *Stack) run() {
	popping := 0

	for {
		select {
		case d := <-s.in:
			s.data = append(s.data, d)
		case <-s.pop:
			popping++
		}

		for popping > 0 && len(s.data) > 0 {
			s.len--
			ret := s.data[len(s.data)-1]
			s.out <- ret
			s.data = s.data[:len(s.data)-1]
			popping--
		}

	}
}

func (s *Stack) Push(d []byte) {
	s.len++
	s.in <- d
}

func (s *Stack) Pop() []byte {
	s.pop <- struct{}{}
	ret := <-s.out
	return ret
}

func (s *Stack) Len() int {
	return s.len
}
