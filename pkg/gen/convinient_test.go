package keygen_test

import (
	"math/rand"
	"testing"
	"time"

	keygen "github.com/onee-only/keygen/pkg/gen"
	"github.com/stretchr/testify/assert"
)

func TestConvGenerator(t *testing.T) {
	gen := keygen.NewConvGenerator(&keygen.ConvConfig{
		BaseConfig: keygen.BaseConfig{
			RandSource: rand.NewSource(time.Now().UnixMicro()),
		},
		Len:       10,
		UseUpper:  true,
		UseLower:  true,
		UseNumber: true,
		UseSymbol: false,
	})

	stream, cancel := gen.GenerateStream()
	defer cancel()

	key := <-stream

	assert.Len(t, key, 10)
}

func BenchmarkConvGenerator(b *testing.B) {
	gen := keygen.NewConvGenerator(&keygen.ConvConfig{
		BaseConfig: keygen.BaseConfig{
			RandSource: rand.NewSource(time.Now().UnixMicro()),
		},
		Len:       10,
		UseUpper:  true,
		UseLower:  true,
		UseNumber: true,
		UseSymbol: false,
	})

	stream, cancel := gen.GenerateStream()
	defer cancel()

	for i := 0; i < b.N; i++ {
		<-stream
	}
}
