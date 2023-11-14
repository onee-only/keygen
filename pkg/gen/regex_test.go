package keygen_test

import (
	"math/rand"
	"regexp"
	"testing"
	"time"

	keygen "github.com/onee-only/keygen/pkg/gen"
	"github.com/stretchr/testify/assert"
)

func TestRegexGenerator(t *testing.T) {
	r := `[0-9]{3}-[0-9]{2}-[0-9]{4}`

	gen, err := keygen.NewRegexGenerator(&keygen.RegexConfig{
		BaseConfig: keygen.BaseConfig{
			RandSource: rand.NewSource(time.Now().UnixMicro()),
		},
		Regex: r,
	})

	if !assert.NoError(t, err) {
		return
	}

	stream, cancel := gen.GenerateStream()
	defer cancel()

	key := <-stream

	matches, err := regexp.Match(r, key)
	assert.NoError(t, err)
	assert.True(t, matches)
}

func BenchmarkRegexGenerator(b *testing.B) {
	r := `[0-9]{3}-[0-9]{2}-[0-9]{4}`

	gen, err := keygen.NewRegexGenerator(&keygen.RegexConfig{
		BaseConfig: keygen.BaseConfig{
			RandSource: rand.NewSource(time.Now().UnixMicro()),
		},
		Regex: r,
	})

	if !assert.NoError(b, err) {
		return
	}

	stream, cancel := gen.GenerateStream()
	defer cancel()

	for i := 0; i < b.N; i++ {
		<-stream
	}
}
