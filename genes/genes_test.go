package genes

import (
	"testing"

	"github.com/paul-hoehne/ea/alleles"
)

func TestGeneGeneration(t *testing.T) {
	gf := GeneFactory{
		AlleleFactories: []alleles.Factory{
			alleles.BitFactory{Width: 4},
			alleles.CodeFactory{Codes: []byte{1, 2, 3}},
		},
	}

	g := gf.Random()

	if len(g.Alleles) != 2 {
		t.Errorf("Expected 2 alleles but got %d", len(g.Alleles))
	}

	if _, ok := g.Alleles[0].(alleles.BitAllele); !ok {
		t.Error("Expected the first allele to be a bitstring")
	}

	if _, ok := g.Alleles[1].(alleles.CodeAllele); !ok {
		t.Error("Expected the second allele to be a coded allele")
	}
}

func TestGeneMutation(t *testing.T) {
	gf := GeneFactory{
		AlleleFactories: []alleles.Factory{
			alleles.BitFactory{Width: 4},
			alleles.NormalRealFactory{Mean: 0, Deviation: 2.0},
		},
	}

	gm := GeneMutator{
		AlleleMutators: []alleles.Mutator{
			alleles.BitMutator{MutationRate: 0.5},
			alleles.NormalRealMutator{Mean: 0, Deviation: 2.0},
		},
	}

	g := gf.Random()
	gprime := gm.Mutate(g)

	if len(gprime.Alleles) != 2 {
		t.Errorf("Expected 2 alleles but got: %d", len(gprime.Alleles))
	}

	if _, ok := g.Alleles[0].(alleles.BitAllele); !ok {
		t.Error("Expected the first allele to be a bit allele")
	}

	if _, ok := g.Alleles[1].(alleles.RealAllele); !ok {
		t.Error("Expected the second allele to be a real allele")
	}
}
