package alleles

import "math/rand"

var masks = []byte{
	byte(1),
	byte(2),
	byte(4),
	byte(8),
	byte(16),
	byte(32),
	byte(64),
	byte(128),
}

// BitAllele is an allele made up strings of bits.  Bit 0 is the low
// order bit of the first byte.  Bit 9 is the lower bit of the second
// byte, etc.  The number of bytes is the ciel(width % 8)
type BitAllele struct {
	Bits  []byte
	Width int
}

// BitFactory produces random strings of bit alleles of the given
// length.  The default frequency of 1's is 0.5, but can be set to
// something else.
type BitFactory struct {
	OnFrequency float64
	Width       int
}

func minInt(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

// Random returns a randomly generated allele
func (b BitFactory) Random() Allele {
	byteWidth := b.Width / 8
	if b.Width%8 != 0 {
		byteWidth++
	}

	result := BitAllele{
		make([]byte, byteWidth),
		b.Width,
	}

	freq := b.OnFrequency
	if freq == 0.0 {
		freq = 0.5
	}

	for i := 0; i < b.Width/8+1; i++ {
		for j := 0; j < minInt(8, b.Width-i*8); j++ {
			if rand.Float64() < freq {
				result.Bits[i] |= masks[j]
			}
		}
	}
	return result
}

// BitMutator mutates a bit in the bit allele according to the mutation
// rate, which should be reasonable (like 1/width)
type BitMutator struct {
	MutationRate float64
}

// Mutate a bit allele to produce another bit allele with some toggled
// bits.  (a 1 becomes a 0, and a 0 becomes a 1).
func (b BitMutator) Mutate(a Allele) Allele {
	bits := a.(BitAllele)
	result := BitAllele{
		Width: bits.Width,
		Bits:  make([]byte, len(bits.Bits)),
	}

	for i := range bits.Bits {
		result.Bits[i] = bits.Bits[i]
	}

	for i := 0; i < result.Width; i++ {
		for j := 0; j < minInt(8, result.Width-i*8); j++ {
			if rand.Float64() < b.MutationRate {
				result.Bits[i] = result.Bits[i] ^ masks[j]
			}
		}
	}

	return result
}
