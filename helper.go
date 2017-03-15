package main

import (
	"strconv"
	"strings"
)

func formatBinary(n uint64, bitCount uint64) string {
	s := strconv.FormatUint(n, 2)
	return strings.Repeat("0", int(bitCount)-len(s)) + s
}
