package keygen

import "math/rand"

type BaseConfig struct {
	RandSource rand.Source
}

func defaultBaseConfig() BaseConfig {
	return BaseConfig{
		RandSource: nil,
	}
}

// Generate generates stream of random strings
type Generator interface {
	GenerateStream() (s <-chan []byte, cancel func())
}
