package infchan

// Channel defines the methods that an unbounded channel must expose.
type Channel[T any] interface {
	In() chan<- T
	Out() <-chan T
	Close()
	Len() int
}

// NewChannel creates a new unbounded channel.  It uses a *DefaultChannel[T]
// underneath.  This is the default implementation.
func NewChannel[T any]() Channel[T] {
	return NewDefaultChannel[T]()
}
