//go:build ignore

package main

import (
	"fmt"

	"github.com/gohryt/zydis-go"
)

func main() {
	request := zydis.EncoderRequest{
		MachineMode:  zydis.MACHINE_MODE_LONG_64,
		Mnemonic:     zydis.MNEMONIC_MOV,
		OperandCount: 2,
		Operands: [5]zydis.EncoderOperand{
			{
				Type: zydis.OPERAND_TYPE_REGISTER,
				Reg:  zydis.EncoderOperandReg{Value: zydis.REGISTER_RAX},
			},
			{
				Type: zydis.OPERAND_TYPE_IMMEDIATE,
				Imm:  zydis.Signed(0x1337),
			},
		},
	}

	buffer, err := zydis.AppendRequest(nil, &request)
	if err != nil {
		panic(err)
	}

	for i := range buffer {
		fmt.Printf("%02X ", buffer[i])
	}

	fmt.Println()
}
