//go:build cgo

package zydis

// #cgo LDFLAGS: -lZydis
// #include <Zydis/Zydis.h>
import "C"
import "unsafe"

func EncoderEncodeInstruction(request, buffer, length unsafe.Pointer) uint32 {
	return uint32(C.ZydisEncoderEncodeInstruction((*C.ZydisEncoderRequest)(request), unsafe.Pointer(buffer), (*C.ulong)(length)))
}
