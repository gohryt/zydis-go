//go:build !cgo

package zydis

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

var EncoderEncodeInstruction func(request, buffer, length unsafe.Pointer) uint32

func init() {
	dl, err := purego.Dlopen("/usr/lib/libZydis.so", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	purego.RegisterLibFunc(&EncoderEncodeInstruction, dl, "ZydisEncoderEncodeInstruction")
}
