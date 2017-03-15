package main

import (
	"strconv"
	"strings"
)

func formatBinary(n uint64, bitCount uint64) string {
	s := strconv.FormatUint(n, 2)
	return strings.Repeat("0", int(bitCount)-len(s)) + s
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
