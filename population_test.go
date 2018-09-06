package ea

import (
	"testing"

	"github.com/paul-hoehne/ea/alleles"
	"github.com/paul-hoehne/ea/genes"
)

func TestPopulationInitialize(t *testing.T) {
	p := Population{
		Size: 10,
		Factory: IndividualFactory{
			GeneFactories: []genes.GeneFactory{
				genes.GeneFactory{
					AlleleFactories: []alleles.Factory{
						alleles.BitFactory{Width: 4},
					},
				},
				genes.GeneFactory{
					AlleleFactories: []alleles.Factory{
						alleles.CodeFactory{Codes: []byte{1, 2, 3}},
					},
				},
			},
		},
	}

	p.Initialize()

	if len(p.Individuals) != 10 {
		t.Errorf("Expected 10 individuals but got: %d", len(p.Individuals))
	}

	if len(p.Individuals[0].Genes) != 2 {
		t.Errorf("Expected individuals to have 2 genes but got: %d",
			len(p.Individuals[0].Genes))
	}

	_, ok := p.Individuals[0].Genes[0].Alleles[0].(alleles.BitAllele)
	if !ok {
		t.Error("Expected first allele to be a bit allele")
	}

	_, ok = p.Individuals[0].Genes[1].Alleles[0].(alleles.CodeAllele)
	if !ok {
		t.Errorf("Expected second allele to be a code allele")
	}
}
