package main

import (
	"github.com/tmthrgd/go-popcount"
)

type Encoder struct {
	ConstraintLength     uint64
	GeneratorPolynomials []uint64
	ShiftRegister        uint64
}

func (enc *Encoder) Flush() {
	enc.ShiftRegister = 0
}

func (enc Encoder) State() uint64 {
	return enc.ShiftRegister >> 1
}

func (enc Encoder) PrintState() {
	formatBinary(enc.State(), enc.ConstraintLength-1)
}

func (enc Encoder) Encode(message uint64, bitCount uint64) uint64 {
	var encoded uint64
	var i uint64

	for i = 0; i < bitCount; i++ {
		encoded <<= uint64(len(enc.GeneratorPolynomials))
		encoded |= enc.encodeBit((message >> i) & 1)
	}

	return encoded
}

func (enc *Encoder) encodeBit(bit uint64) uint64 {
	var word uint64

	enc.ShiftRegister >>= 1
	enc.ShiftRegister |= (bit << (enc.ConstraintLength - 1))

	for _, gen := range enc.GeneratorPolynomials {
		word <<= 1
		word |= (popcount.Count64(enc.ShiftRegister&gen) % 2)
	}

	return word
}
