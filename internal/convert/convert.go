package convert

import "unsafe"

func StrToBytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
