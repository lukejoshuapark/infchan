package infchan

// DefaultChannel is the default implementation of an unbounded channel.
type DefaultChannel[T any] struct {
	in  chan T
	out chan T
	buf []T
}

var _ Channel[struct{}] = &DefaultChannel[struct{}]{}

func NewDefaultChannel[T any]() *DefaultChannel[T] {
	c := &DefaultChannel[T]{
		in:  make(chan T),
		out: make(chan T),
		buf: []T{},
	}

	go c.process()
	return c
}

func (c *DefaultChannel[T]) In() chan<- T {
	return c.in
}

func (c *DefaultChannel[T]) Out() <-chan T {
	return c.out
}

func (c *DefaultChannel[T]) Close() {
	close(c.in)
}

func (c *DefaultChannel[T]) Len() int {
	return len(c.buf)
}

func (c *DefaultChannel[T]) process() {
	for c.in != nil || len(c.buf) > 0 {
		var outValue T
		var outChan chan<- T
		if len(c.buf) > 0 {
			outValue = c.buf[0]
			outChan = c.out
		}

		select {
		case inValue, ok := <-c.in:
			if !ok {
				c.in = nil
				continue
			}
			c.buf = append(c.buf, inValue)

		case outChan <- outValue:
			c.buf = c.buf[1:]
		}
	}

	close(c.out)
	c.out = nil
}
