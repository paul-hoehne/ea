package alleles

import "testing"

func TestSingleByteBitFactory(t *testing.T) {
	b := BitFactory{
		Width:       4,
		OnFrequency: 1.0,
	}

	bits := b.Random().(BitAllele)

	if bits.Width != 4 {
		t.Errorf("Expected width to be 4 but got %d", bits.Width)
	}

	if len(bits.Bits) != 1 {
		t.Errorf("Expected one byte but got %d", len(bits.Bits))
	}

	if bits.Bits[0] != 0x0f {
		t.Errorf("Expected the bottom 4 bits set but %b", bits.Bits[0])
	}
}

func TestBitFactoryNoFrequency(t *testing.T) {
	b := BitFactory{
		Width: 8,
	}

	bits := b.Random().(BitAllele)

	if bits.Bits[0] == 0 {
		t.Errorf("Expected the bits to be non zero: %b", bits.Bits[0])
	}

	if bits.Bits[0] == 0xff {
		t.Errorf("Expected the bits to be all non-one: %b", bits.Bits[0])
	}
}

func TestMultiByteBitFactory(t *testing.T) {
	b := BitFactory{
		Width:       12,
		OnFrequency: 1.0,
	}

	bits := b.Random().(BitAllele)

	if bits.Width != 12 {
		t.Errorf("Expected width to be 12 but got %d", bits.Width)
	}

	if len(bits.Bits) != 2 {
		t.Errorf("Expected two bytes but got %d", len(bits.Bits))
	}

	if bits.Bits[0] != 0xff || bits.Bits[1] != 0x0f {
		t.Errorf("Expected low to be ff but got %b and high to be 0x0f but got %b",
			bits.Bits[0], bits.Bits[1])
	}

	b = BitFactory{
		Width:       24,
		OnFrequency: 1.0,
	}

	bits = b.Random().(BitAllele)

	if len(bits.Bits) != 3 {
		t.Errorf("Expected three bytes but got %d", len(bits.Bits))
	}

	if bits.Bits[2] != 0xff || bits.Bits[1] != 0xff || bits.Bits[0] != 0xff {
		t.Errorf("Expected all 24 bits set but got %b %b %b", bits.Bits[0],
			bits.Bits[1], bits.Bits[2])
	}
}

func TestSingleByteMutate(t *testing.T) {
	b := BitFactory{
		Width:       4,
		OnFrequency: 1.0,
	}

	allele := b.Random()

	m := BitMutator{
		MutationRate: 1.0,
	}

	allele = m.Mutate(allele)

	bits := allele.(BitAllele)

	if len(bits.Bits) != 1 {
		t.Errorf("There should have been one byte but got: %d", len(bits.Bits))
	}

	if bits.Width != 4 {
		t.Errorf("The width should have been 4 but got: %d", b.Width)
	}

	if bits.Bits[0] != 0 {
		t.Errorf("Expected all bits to be flipped but got: %x", bits.Bits[0])
	}
}

func TestMultiByteMutate(t *testing.T) {
	b := BitFactory{
		Width:       12,
		OnFrequency: 1.0,
	}

	allele := b.Random()
	m := BitMutator{
		MutationRate: 1.0,
	}
	allele = m.Mutate(allele)

	bits := allele.(BitAllele)

	if bits.Width != 12 {
		t.Errorf("Expected width was 12 but got %d", bits.Width)
	}

	if len(bits.Bits) != 2 {
		t.Errorf("Expected two bytes but got %d", len(bits.Bits))
	}

	if bits.Bits[0] != 0 || bits.Bits[1] != 0 {
		t.Errorf("Expected zero bytes but got %b %b", bits.Bits[0], bits.Bits[1])
	}

	b = BitFactory{
		Width:       24,
		OnFrequency: 1.0,
	}

	allele = b.Random()
	allele = m.Mutate(allele)

	bits = allele.(BitAllele)

	if len(bits.Bits) != 3 {
		t.Errorf("Expected three bytes but got %d", len(bits.Bits))
	}

	if bits.Bits[0] != 0 || bits.Bits[1] != 0 || bits.Bits[2] != 0 {
		t.Errorf("Expected bits to be zero but got %b %b %b", bits.Bits[0],
			bits.Bits[1], bits.Bits[2])
	}
}

func TestBitStringString(t *testing.T) {
	bs := BitAllele{
		Bits:  []byte{0x0f},
		Width: 4,
	}

	if "1111" != bs.String() {
		t.Errorf("Expected 1111 but got '%s'", bs.String())
	}

	bs.Bits[0] = 7

	if "0111" != bs.String() {
		t.Errorf("Expected 0111 but got '%s'", bs.String())
	}

	bs.Bits = []byte{0x03, 0x43}
	bs.Width = 12

	if "0011 01000011" != bs.String() {
		t.Errorf("Expected 0011 01000011 but got '%s'", bs.String())
	}

	bs.Bits = []byte{0xf3, 0xe1}
	bs.Width = 16

	if "11110011 11100001" != bs.String() {
		t.Errorf("Expected 11110011 11100001 but got '%s'", bs.String())
	}
}

func TestBitAlleleCopy(t *testing.T) {
	bs1 := BitAllele{
		Width: 16,
		Bits:  []byte{0xab, 0xcd},
	}

	copy := bs1.Copy()

	bs2, ok := copy.(BitAllele)
	if !ok {
		t.Error("Expected a type of BitAllele")
	}

	if bs1.Width != bs2.Width {
		t.Errorf("Expected width to be %d but got %d", bs1.Width, bs2.Width)
	}

	bs1.Bits[0] = 0x00

	if bs2.Bits[0] == 0x00 {
		t.Error("Edited original and change showed up in copy")
	}
}
