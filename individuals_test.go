package ea

import (
	"testing"

	"github.com/chilts/sid"

	"github.com/paul-hoehne/ea/alleles"
	"github.com/paul-hoehne/ea/genes"
)

func TestIndividualFactoryCreate(t *testing.T) {
	ifact := IndividualFactory{
		GeneFactories: []genes.GeneFactory{
			genes.GeneFactory{
				AlleleFactories: []alleles.Factory{
					alleles.BitFactory{Width: 4, OnFrequency: 1.0},
				},
			},
			genes.GeneFactory{
				AlleleFactories: []alleles.Factory{
					alleles.BitFactory{Width: 4, OnFrequency: 1.0},
					alleles.CodeFactory{Codes: []byte{2}},
				},
			},
		},
	}

	i := ifact.Create()

	if i.ID == "" {
		t.Error("Expected the individual to have an ID")
	}

	if len(i.Genes) != 2 {
		t.Errorf("Expected 2 genes but got %d", len(i.Genes))
	}

	if len(i.Genes[0].Alleles) != 1 || len(i.Genes[1].Alleles) != 2 {
		t.Errorf("Expected 1 allele in the first gene and 2 in the second but got: %d and %d",
			len(i.Genes[0].Alleles), len(i.Genes[1].Alleles))
	}

	if _, ok := i.Genes[0].Alleles[0].(alleles.BitAllele); !ok {
		t.Error("Expected first allele to be a bit Allele")
	}

	if _, ok := i.Genes[1].Alleles[0].(alleles.BitAllele); !ok {
		t.Error("Expected second allele to be a bit allele")
	}

	if _, ok := i.Genes[1].Alleles[1].(alleles.CodeAllele); !ok {
		t.Error("Expected third allele to be a code allele")
	}
}

func TestIndividualSpawn(t *testing.T) {
	i := Individual{
		ID: sid.Id(),
		Genes: []genes.Gene{
			genes.Gene{
				Alleles: []alleles.Allele{
					alleles.BitAllele{Width: 4, Bits: []byte{1}},
				},
			},
			genes.Gene{
				Alleles: []alleles.Allele{
					alleles.RealAllele(0.5),
				},
			},
			genes.Gene{
				Alleles: []alleles.Allele{
					alleles.CodeAllele{Value: 1},
				},
			},
			genes.Gene{
				Alleles: []alleles.Allele{
					alleles.RealAllele(0.1),
					alleles.BitAllele{Width: 6, Bits: []byte{5}},
					alleles.CodeAllele{Value: 7},
				},
			},
		},
	}

	ix := i.Spawn()

	if ix.ID == "" || ix.ID == i.ID {
		t.Error("Expected spawned child to have new, unique ID")
	}

	if len(ix.Genes) != 4 {
		t.Errorf("Expected 4 genes but got %d", len(ix.Genes))
	}

	if len(ix.Genes[0].Alleles) != 1 || len(ix.Genes[1].Alleles) != 1 ||
		len(ix.Genes[2].Alleles) != 1 || len(ix.Genes[3].Alleles) != 3 {
		t.Errorf("Expected 1, 1, 1, and 3 alleles but got %d, %d, %d and %d",
			len(ix.Genes[0].Alleles), len(ix.Genes[1].Alleles),
			len(ix.Genes[2].Alleles), len(ix.Genes[3].Alleles))
	}

	if _, ok := ix.Genes[0].Alleles[0].(alleles.BitAllele); !ok {
		t.Error("Expected first allele to be a bit allele")
	}

	if _, ok := ix.Genes[1].Alleles[0].(alleles.RealAllele); !ok {
		t.Error("Expected second allele to be a real allele")
	}

	if _, ok := ix.Genes[2].Alleles[0].(alleles.CodeAllele); !ok {
		t.Errorf("Expected third allele to be a code allele")
	}

	if _, ok := ix.Genes[3].Alleles[0].(alleles.RealAllele); !ok {
		t.Error("Expected fourth allele to be a real allele")
	}

	if _, ok := ix.Genes[3].Alleles[1].(alleles.BitAllele); !ok {
		t.Error("Expected fifth allele to be a bit allele")
	}

	if _, ok := ix.Genes[3].Alleles[2].(alleles.CodeAllele); !ok {
		t.Errorf("Expected sixth allele to be a code allele")
	}
}

func TestIndividualBreed(t *testing.T) {
	p1 := Individual{
		ID: sid.Id(),
		Genes: []genes.Gene{
			genes.Gene{
				Alleles: []alleles.Allele{
					alleles.BitAllele{Width: 4, Bits: []byte{1}},
				},
			},
			genes.Gene{
				Alleles: []alleles.Allele{
					alleles.CodeAllele{Value: 4},
				},
			},
		},
	}

	p2 := Individual{
		ID: sid.Id(),
		Genes: []genes.Gene{
			genes.Gene{
				Alleles: []alleles.Allele{
					alleles.BitAllele{Width: 4, Bits: []byte{10}},
				},
			},
			genes.Gene{
				Alleles: []alleles.Allele{
					alleles.CodeAllele{Value: 1},
				},
			},
		},
	}

	p3 := p1.Breed(p2, 1.0)

	if p3.ID == "" || p3.ID == p1.ID || p3.ID == p2.ID {
		t.Error("Expected child to have a unique ID")
	}

	if len(p3.Genes) != 2 {
		t.Errorf("Expected child to have 2 genes but got: %d", len(p3.Genes))
	}

	if len(p3.Genes[0].Alleles) != 1 || len(p3.Genes[1].Alleles) != 1 {
		t.Errorf("Expected each gene to have 1 allele but got %d and %d",
			len(p3.Genes[0].Alleles), len(p3.Genes[0].Alleles))
	}

	b, ok := p3.Genes[0].Alleles[0].(alleles.BitAllele)
	if !ok {
		t.Errorf("Expected the first alleles of the child to be a bit allele")
	}

	if b.Bits[0] != 10 {
		t.Errorf("Expected crossover value to be 10 but got %d", b.Bits[0])
	}

	c, ok := p3.Genes[1].Alleles[0].(alleles.CodeAllele)
	if !ok {
		t.Errorf("Expected the second allele of the child to be a code allele")
	}

	if c.Value != 1 {
		t.Errorf("Expected the second crossover value to be 1 but got: %d", c.Value)
	}

	p3 = p1.Breed(p2, 0.0)

	b, ok = p3.Genes[0].Alleles[0].(alleles.BitAllele)
	if b.Bits[0] != 1 {
		t.Errorf("Expected crossover value to be 1 but got %d", b.Bits[0])
	}

	c, ok = p3.Genes[1].Alleles[0].(alleles.CodeAllele)
	if c.Value != 4 {
		t.Errorf("Expected crossover value to be 4 but got %d", c.Value)
	}
}

func TestIndividualMutate(t *testing.T) {
	p1 := Individual{
		ID: sid.Id(),
		Genes: []genes.Gene{
			genes.Gene{
				Alleles: []alleles.Allele{
					alleles.BitAllele{Width: 4, Bits: []byte{8}},
				},
			},
			genes.Gene{
				Alleles: []alleles.Allele{
					alleles.CodeAllele{Value: 4},
				},
			},
		},
	}

	ms := []genes.GeneMutator{
		genes.GeneMutator{
			AlleleMutators: []alleles.Mutator{
				alleles.BitMutator{MutationRate: 1.0},
			},
		},
		genes.GeneMutator{
			AlleleMutators: []alleles.Mutator{
				alleles.CodeMutator{Codes: []byte{1, 2, 3}},
			},
		},
	}

	p1.Mutate(ms, 1.0)

	if p1.ID == "" {
		t.Error("Expected mutated individual to have an ID")
	}

	b, ok := p1.Genes[0].Alleles[0].(alleles.BitAllele)
	if !ok {
		t.Error("Expected first allele to be a bit allele")
	}

	if b.Bits[0] != 0x7 {
		t.Errorf("Expected mutated bits to be 0x0c but got %x", b.Bits[0])
	}

	c, ok := p1.Genes[1].Alleles[0].(alleles.CodeAllele)
	if !ok {
		t.Error("Expected second allele to be a code allele")
	}

	if c.Value == 4 {
		t.Error("Expected mutated value to not be 4")
	}

	p1 = Individual{
		Genes: []genes.Gene{
			genes.Gene{
				Alleles: []alleles.Allele{
					alleles.BitAllele{Width: 4, Bits: []byte{8}},
				},
			},
			genes.Gene{
				Alleles: []alleles.Allele{
					alleles.CodeAllele{Value: 4},
				},
			},
		},
	}

	p1.Mutate(ms, 0.0)
	b = p1.Genes[0].Alleles[0].(alleles.BitAllele)
	if b.Bits[0] != 8 {
		t.Errorf("Expected mutated bits to be 8, but got %x", b.Bits[0])
	}

	c = p1.Genes[1].Alleles[0].(alleles.CodeAllele)
	if c.Value != 4 {
		t.Errorf("Expected value to be four but got %d", c.Value)
	}
}
