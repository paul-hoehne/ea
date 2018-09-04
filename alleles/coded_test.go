package alleles

import "testing"

func TestRandomCodedAllele(t *testing.T) {
	f := CodeFactory{
		Codes:       []byte{1, 2, 3, 4},
		Frequencies: []float64{0.0, 1.0, 0.0, 0.0},
	}

	a := f.Random().(CodeAllele)

	if a.Value != 2 {
		t.Errorf("Expected value to be 2 but was: %d", a.Value)
	}
}

func TestRandomCodedAlleleNoFrequences(t *testing.T) {
	f := CodeFactory{
		Codes: []byte{1, 2, 3},
	}

	a := f.Random()
	c := a.(CodeAllele)

	if c.Value != 1 && c.Value != 2 && c.Value != 3 {
		t.Errorf("Should have been 1, 2, or 3: %d", c.Value)
	}
}

func TestRandomCodedMutator(t *testing.T) {
	f := CodeFactory{
		Codes:       []byte{1, 2},
		Frequencies: []float64{0.0, 1.0},
	}

	m := CodeMutator{
		Codes:       []byte{1, 2},
		Frequencies: []float64{1.0, 0.0},
	}

	a := f.Random()
	aprime := m.Mutate(a)

	c := aprime.(CodeAllele)

	if c.Value != 1 {
		t.Errorf("Expected 1 but got: %d", c.Value)
	}
}
