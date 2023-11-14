package keygen

import (
	"bytes"
	"math/rand"
)

var (
	lowerCharset  = "abcdefghijklnmopqrstuvwxyz"
	upperCharset  = "ABCDEFGHIJKLNMOPQRSTUVWXYZ"
	numberCharset = "1234567890"
	symbolCharset = "`!@#$%^&*()_-+={[}]|\\:;\"'<,>.?/"
)

// ConvConfig is configuration for generating random string with convinience.
type ConvConfig struct {
	BaseConfig
	Len       uint64
	UseUpper  bool
	UseLower  bool
	UseNumber bool
	UseSymbol bool
}

func DefaultConvConfig() *ConvConfig {
	return &ConvConfig{
		BaseConfig: defaultBaseConfig(),
		UseUpper:   true,
		UseLower:   true,
		UseNumber:  true,
		UseSymbol:  false,
	}
}

func NewConvGenerator(conf *ConvConfig) Generator {
	gen := customGenerator{
		conf: &CustomConfig{
			BaseConfig: conf.BaseConfig,
			Len:        conf.Len,
		},
	}

	buf := bytes.NewBuffer(make([]byte, 0))
	if conf.UseLower {
		buf.WriteString(lowerCharset)
	}
	if conf.UseUpper {
		buf.WriteString(upperCharset)
	}
	if conf.UseNumber {
		buf.WriteString(numberCharset)
	}
	if conf.UseSymbol {
		buf.WriteString(symbolCharset)
	}

	gen.charset = buf.Bytes()
	gen.rand = rand.New(conf.RandSource)

	return &gen
}
