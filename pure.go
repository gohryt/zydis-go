package zydis

import (
	"github.com/ebitengine/purego"
)

var EncoderEncodeInstruction func(request *EncoderRequest, buffer *byte, length *uint) uint32

func init() {
	dl, err := purego.Dlopen("/usr/lib/libZydis.so", purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	purego.RegisterLibFunc(&EncoderEncodeInstruction, dl, "ZydisEncoderEncodeInstruction")
}
