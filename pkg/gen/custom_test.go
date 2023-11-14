package keygen_test

import (
	"math/rand"
	"testing"
	"time"

	keygen "github.com/onee-only/keygen/pkg/gen"
	"github.com/stretchr/testify/assert"
)

func TestCustomGenerator(t *testing.T) {
	charset := []byte("abcdef")

	gen := keygen.NewCustomGenerator(&keygen.CustomConfig{
		BaseConfig: keygen.BaseConfig{
			RandSource: rand.NewSource(time.Now().UnixMicro()),
		},
		Len:      10,
		Chartset: charset,
	})

	stream, cancel := gen.GenerateStream()
	defer cancel()

	key := <-stream

	assert.Len(t, key, 10)
}

func BenchmarkCustomGenerator(b *testing.B) {
	charset := []byte("abcdef")

	gen := keygen.NewCustomGenerator(&keygen.CustomConfig{
		BaseConfig: keygen.BaseConfig{
			RandSource: rand.NewSource(time.Now().UnixMicro()),
		},
		Len:      10,
		Chartset: charset,
	})

	stream, cancel := gen.GenerateStream()
	defer cancel()

	for i := 0; i < b.N; i++ {
		<-stream
	}
}
