![](icon.png)

# infchan

Exposes a generic, unbounded channel.

## Usage Example

```go
c := infchan.NewChannel[int]()

for i := 0; i < 100; i++ {
    c.In() <- i
}
c.Close()

sum := 0
for n := range c.Out() {
    sum += n
}

test.That(t, sum, is.EqualTo(4950))
```

## Notes

When the consuming side of a channel is processing slower than the producing
side, bad things happen.  These bad things are unavoidable, but we can decide
what flavor of "bad thing" we'd like to deal with!

The native buffered channel in Go will cause sends to block when the buffer is
full.

The channel exposed in this implementation will continously accept and buffer
sends until the application runs out of memory.

Carefully consider which one of these you'd prefer - you may find that your
problem actually lies elsewhere.
