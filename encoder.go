package zydis

import (
	"fmt"
	"io"
	"unsafe"
)

type (
	EncoderRequest struct {
		MachineMode      MachineMode
		AllowedEncodings EncodableEncoding
		Mnemonic         Mnemonic
		Prefixes         InstructionAttributes
		BranchType       BranchType
		BranchWidth      BranchWidth
		AddressSizeHint  AddressSizeHint
		OperandSizeHint  OperandSizeHint
		OperandCount     uint8
		Operands         [5]EncoderOperand
		Evex             EncoderRequestEvexFeatures
		Mvex             EncoderRequestMvexFeatures
	}

	EncoderRequestEvexFeatures struct {
		Broadcast   BroadcastMode
		Rounding    RoundingMode
		Sae         Bool
		ZeroingMask Bool
	}

	EncoderRequestMvexFeatures struct {
		Broadcast    BroadcastMode
		Conversion   ConversionMode
		Rounding     RoundingMode
		Swizzle      SwizzleMode
		Sae          Bool
		EvictionHint Bool
	}
)

type Encoder struct {
	writer io.Writer
	buffer [MAX_INSTRUCTION_LENGTH]byte
}

func NewEncoder(writer io.Writer) *Encoder {
	return &Encoder{
		writer: writer,
	}
}

func (encoder *Encoder) Encode(request *EncoderRequest) error {
	length := MAX_INSTRUCTION_LENGTH

	status := EncoderEncodeInstruction(request, &encoder.buffer[0], &length)
	if !Success(status) {
		return fmt.Errorf("%w: %d", ErrWrongStatus, status)
	}

	_, err := encoder.writer.Write(encoder.buffer[:length])
	if err != nil {
		return err
	}

	return nil
}

func AppendRequest(to []byte, request *EncoderRequest) ([]byte, error) {
	length := MAX_INSTRUCTION_LENGTH
	buffer := make([]byte, length)

	status := EncoderEncodeInstruction(request, unsafe.SliceData(buffer), &length)
	if !Success(status) {
		return nil, fmt.Errorf("%w: %d", ErrWrongStatus, status)
	}

	return append(to, buffer[:length]...), nil
}
