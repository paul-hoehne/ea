package alleles

import "testing"

func TestRandomRealAlleleFactory(t *testing.T) {
	u := UniformBoundedRealFactory{
		-1.0,
		1.0,
	}

	for i := 0; i < 10; i++ {
		ra := u.Random()

		if ra == 0.0 {
			t.Errorf("Expected non zero")
		}

		if ra.(RealAllele) > u.Max || ra.(RealAllele) < u.Min {
			t.Errorf("Out of range: %v", ra.(RealAllele))
			break
		}
	}
}

func TestUniformRealMutator(t *testing.T) {
	u := UinformBoundedRealMutator{
		1.0,
		2.0,
	}

	var ra RealAllele

	mutated := u.Mutate(ra)
	if mutated.(RealAllele) < u.Min || mutated.(RealAllele) > u.Max {
		t.Errorf("Mutated value out of range: %v", mutated)
	}
}

func TestNormalRealAlleleFactory(t *testing.T) {
	n := NormalRealFactory{
		Mean:      1.0,
		Deviation: 0.1,
	}

	a := n.Random()

	// Note this can fail if the value is outside of 6 standard deviations.
	if a.(RealAllele) > RealAllele(1.6) || a.(RealAllele) < RealAllele(0.4) {
		t.Errorf("Expected a value around 1.0 but got %v", a)
	}
}

func TestNormalRealAlleleMutator(t *testing.T) {
	n := NormalRealMutator{
		Mean:      1.0,
		Deviation: 0.1,
	}

	a := RealAllele(0.0)
	mut := n.Mutate(a)

	if mut.(RealAllele) > RealAllele(1.6) || mut.(RealAllele) < RealAllele(0.4) {
		t.Errorf("Expected a value around 1.0 but got %v", mut)
	}
}
