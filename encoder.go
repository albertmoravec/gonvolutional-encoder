package main

import (
	"github.com/tmthrgd/go-popcount"
)

type Encoder struct {
	ConstraintLength     uint64
	GeneratorPolynomials []uint64
	Register             uint64
}

func (enc *Encoder) Flush() {
	enc.Register = 0
}

func (enc Encoder) State() uint64 {
	return enc.Register >> 1
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

	enc.shiftRegister()
	enc.insertBit(bit)

	for _, gen := range enc.GeneratorPolynomials {	// TODO Optimize for identical polynomials
		word <<= 1
		word |= enc.calculateOutput(gen)
	}

	return word
}

func (enc *Encoder) shiftRegister() {
	enc.Register >>= 1
}

func (enc *Encoder) insertBit(bit uint64) {
	enc.Register |= (bit << (enc.ConstraintLength - 1))
}

func (enc Encoder) calculateOutput(polynomial uint64) uint64 {
	return popcount.Count64(enc.Register & polynomial) % 2
}
