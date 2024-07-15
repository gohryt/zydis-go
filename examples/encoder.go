//go:build ignore

package main

import (
	"bytes"
	"fmt"
	"syscall"
	"unsafe"

	"github.com/gohryt/zydis-go"
)

const example = "Hello world!\n"

func main() {
	buffer := new(bytes.Buffer)

	encoder := zydis.NewEncoder(buffer)

	err := encoder.Encode(&zydis.EncoderRequest{
		MachineMode:  zydis.MACHINE_MODE_LONG_64,
		Mnemonic:     zydis.MNEMONIC_MOV,
		OperandCount: 2,
		Operands: [5]zydis.EncoderOperand{
			{Type: zydis.OPERAND_TYPE_REGISTER, Reg: zydis.EncoderOperandReg{Value: zydis.REGISTER_RAX}},
			{Type: zydis.OPERAND_TYPE_IMMEDIATE, Imm: zydis.Signed(0x1)},
		},
	})
	if err != nil {
		panic(err)
	}

	err = encoder.Encode(&zydis.EncoderRequest{
		MachineMode:  zydis.MACHINE_MODE_LONG_64,
		Mnemonic:     zydis.MNEMONIC_MOV,
		OperandCount: 2,
		Operands: [5]zydis.EncoderOperand{
			{Type: zydis.OPERAND_TYPE_REGISTER, Reg: zydis.EncoderOperandReg{Value: zydis.REGISTER_RDI}},
			{Type: zydis.OPERAND_TYPE_IMMEDIATE, Imm: zydis.Signed(0x1)},
		},
	})
	if err != nil {
		panic(err)
	}

	err = encoder.Encode(&zydis.EncoderRequest{
		MachineMode:  zydis.MACHINE_MODE_LONG_64,
		Mnemonic:     zydis.MNEMONIC_MOV,
		OperandCount: 2,
		Operands: [5]zydis.EncoderOperand{
			{Type: zydis.OPERAND_TYPE_REGISTER, Reg: zydis.EncoderOperandReg{Value: zydis.REGISTER_RDX}},
			{Type: zydis.OPERAND_TYPE_IMMEDIATE, Imm: zydis.Signed(int64(len(example)))},
		},
	})
	if err != nil {
		panic(err)
	}

	err = encoder.Encode(&zydis.EncoderRequest{
		MachineMode:  zydis.MACHINE_MODE_LONG_64,
		Mnemonic:     zydis.MNEMONIC_MOV,
		OperandCount: 2,
		Operands: [5]zydis.EncoderOperand{
			{Type: zydis.OPERAND_TYPE_REGISTER, Reg: zydis.EncoderOperandReg{Value: zydis.REGISTER_RSI}},
			{Type: zydis.OPERAND_TYPE_IMMEDIATE, Imm: zydis.Unsigned(uint64(uintptr(unsafe.Pointer(unsafe.StringData(example)))))},
		},
	})
	if err != nil {
		panic(err)
	}

	err = encoder.Encode(&zydis.EncoderRequest{
		MachineMode: zydis.MACHINE_MODE_LONG_64,
		Mnemonic:    zydis.MNEMONIC_SYSCALL,
	})
	if err != nil {
		panic(err)
	}

	err = encoder.Encode(&zydis.EncoderRequest{
		MachineMode: zydis.MACHINE_MODE_LONG_64,
		Mnemonic:    zydis.MNEMONIC_RET,
	})
	if err != nil {
		panic(err)
	}

	bytes := buffer.Bytes()

	for _, b := range bytes {
		fmt.Printf("%02X ", b)
	}
	fmt.Println()

	executable, err := syscall.Mmap(-1, 0, len(bytes), syscall.PROT_READ|syscall.PROT_WRITE|syscall.PROT_EXEC, syscall.MAP_PRIVATE|syscall.MAP_ANONYMOUS)
	if err != nil {
		panic(err)
	}
	copy(executable, bytes)

	unsafeExecutable := (uintptr)(unsafe.Pointer(&executable))
	print := *(*func())(unsafe.Pointer(&unsafeExecutable))
	print()
}
