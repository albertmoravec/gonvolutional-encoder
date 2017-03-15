package main

import (
	"flag"
	"fmt"
)

var (
	Constraint       uint64
	InputPolynomials []uint64
	Input            uint64
	BitCount         uint64

	EncoderObject Encoder
)

func init() {
	flag.Uint64Var(&Constraint, "k", 0, "Constraint length")
	polynomials := flag.String("g", "", "Generator polynomials")	// TODO throw error on 0 polynomial
	input := flag.String("i", "", "Input message")
	flag.Parse()

	InputPolynomials = splitPolynomials(*polynomials)
	Input, BitCount = transformMessage(*input)
}

func main() {
	EncoderObject = Encoder{	// TODO Create constructor func
		ConstraintLength:     Constraint,
		GeneratorPolynomials: InputPolynomials,
		Register:             0,
	}

	fmt.Println(formatBinary(EncoderObject.Encode(Input, BitCount), BitCount*uint64(len(InputPolynomials))))
}
