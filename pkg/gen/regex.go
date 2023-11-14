package keygen

import (
	"github.com/onee-only/keygen/internal/convert"
	"github.com/pkg/errors"
	regen "github.com/zach-klippenstein/goregen"
)

// Config is configuration for regex matching random string.
type RegexConfig struct {
	BaseConfig
	Regex string
}

type regexGenerator struct {
	gen regen.Generator
}

func NewRegexGenerator(conf *RegexConfig) (Generator, error) {
	gen := regexGenerator{}

	var args regen.GeneratorArgs

	args.RngSource = conf.RandSource

	regen, err := regen.NewGenerator(conf.Regex, &args)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create generator")
	}

	gen.gen = regen

	return &gen, nil
}

func (g *regexGenerator) GenerateStream() (s <-chan []byte, cancel func()) {
	done := make(chan struct{})
	stream := make(chan []byte)

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- convert.StrToBytes(g.gen.Generate()):
			}
		}
	}()
	return stream, func() { done <- struct{}{}; close(done) }
}
