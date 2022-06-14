package infchan

import (
	"testing"

	"github.com/lukejoshuapark/test"
	"github.com/lukejoshuapark/test/is"
)

func TestAllowSynchronousWriteThenRead(t *testing.T) {
	c := NewDefaultChannel[int]()

	for i := 0; i < 100; i++ {
		c.In() <- i
	}
	c.Close()

	sum := 0
	for n := range c.Out() {
		sum += n
	}

	test.That(t, sum, is.EqualTo(4950))
}

func TestAllowAsynchronousWriteAndRead(t *testing.T) {
	c := NewDefaultChannel[int]()

	go func() {
		for i := 0; i < 100; i++ {
			c.In() <- i
		}
		c.Close()
	}()

	sum := 0
	for n := range c.Out() {
		sum += n
	}

	test.That(t, sum, is.EqualTo(4950))
}
