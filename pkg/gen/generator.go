package keygen

import (
	"math/rand"
	"time"
)

type BaseConfig struct {
	RandSource rand.Source
}

func DefaultBaseConfig() BaseConfig {
	return BaseConfig{
		RandSource: rand.NewSource(time.Now().UnixNano()),
	}
}

// Generate generates stream of random strings
type Generator interface {
	GenerateStream() (s <-chan []byte, cancel func())
}
