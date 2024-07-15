package zydis

import (
	"errors"
	"unsafe"
)

type (
	Bool                  uint8
	InstructionAttributes uint64

	Imm struct {
		raw [8]uint8
	}
)

func Unsigned(unsigned uint64) (imm Imm) {
	imm.SetUnsigned(unsigned)
	return imm
}

func Signed(signed int64) (imm Imm) {
	imm.SetSigned(signed)
	return imm
}

func (imm *Imm) SetUnsigned(unsigned uint64) {
	*(*uint64)(unsafe.Pointer(imm)) = unsigned
}

func (imm *Imm) GetUnsigned() uint64 {
	return *(*uint64)(unsafe.Pointer(imm))
}

func (imm *Imm) SetSigned(signed int64) {
	*(*int64)(unsafe.Pointer(imm)) = signed
}

func (imm *Imm) GetSigned() int64 {
	return *(*int64)(unsafe.Pointer(imm))
}

type (
	EncoderOperand struct {
		Type OperandType
		Reg  EncoderOperandReg
		Mem  EncoderOperandMem
		Ptr  EncoderOperandPtr
		Imm  EncoderOperandImm
	}

	EncoderOperandReg struct {
		Value Register
		Is4   Bool
	}

	EncoderOperandMem struct {
		Base         Register
		Index        Register
		Scale        uint8
		Displacement int64
		Size         uint16
	}

	EncoderOperandPtr struct {
		Segment uint16
		Offset  uint32
	}

	EncoderOperandImm = Imm
)

func Success(status uint32) bool {
	return int32(status) >= 0
}

func Failed(status uint32) bool {
	return int32(status) < 0
}

const (
	MAX_INSTRUCTION_LENGTH uint = 15
)

var ErrWrongStatus = errors.New("wrong status")
