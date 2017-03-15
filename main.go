package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
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
	polynomials := flag.String("g", "", "Generator polynomials")
	input := flag.String("i", "", "Input message")
	flag.Parse()

	InputPolynomials = splitPolynomials(*polynomials)
	Input, BitCount = transformMessage(*input)
}

func main() {
	EncoderObject = Encoder{
		ConstraintLength:     Constraint,
		GeneratorPolynomials: InputPolynomials,
		ShiftRegister:        0,
	}

	fmt.Println(strconv.FormatUint(EncoderObject.Encode(Input, BitCount), 2))

}

func splitPolynomials(input string) []uint64 {
	var polynomials []uint64

	for _, polynomial := range strings.Split(input, " ") {
		parsed, err := strconv.ParseUint(polynomial, 8, 32)
		if err != nil {
			return nil
		}

		polynomials = append(polynomials, parsed)
	}

	return polynomials
}

func transformMessage(message string) (uint64, uint64) {
	parsed, err := strconv.ParseUint(message, 2, 64)
	if err != nil {
		return 0, 0
	}

	return parsed, uint64(len(message))
}
