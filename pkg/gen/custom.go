package keygen

import (
	"math/rand"
)

// Config is configuration for generating random string.
type CustomConfig struct {
	BaseConfig
	Len      uint64
	Chartset []byte
}

type customGenerator struct {
	conf    *CustomConfig
	rand    *rand.Rand
	charset []byte
}

func NewCustomGenerator(conf *CustomConfig) Generator {
	gen := customGenerator{conf: conf}

	gen.charset = conf.Chartset
	gen.rand = rand.New(conf.RandSource)

	return &gen
}

func (g *customGenerator) GenerateStream() (s <-chan []byte, cancel func()) {
	done := make(chan struct{})
	stream := make(chan []byte)

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- g.makeNew():
			}
		}
	}()
	return stream, func() { done <- struct{}{}; close(done) }
}

func (g *customGenerator) makeNew() []byte {

	buf := make([]byte, g.conf.Len)
	for i := 0; i < int(g.conf.Len); i++ {
		idx := g.rand.Intn(len(g.charset))
		buf[i] = g.charset[idx]
	}

	return buf
}
