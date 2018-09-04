package genes

import "github.com/paul-hoehne/ea/alleles"

// Gene is collections of alleles that can be expressed
// as a value.
type Gene struct {
	Alleles []alleles.Allele
}

// GeneFactory produces a new random gene from the given alleles
type GeneFactory struct {
	AlleleFactories []alleles.Factory
}

// Random creates a random gene
func (gf GeneFactory) Random() Gene {
	result := Gene{
		make([]alleles.Allele, len(gf.AlleleFactories)),
	}

	for i, f := range gf.AlleleFactories {
		result.Alleles[i] = f.Random()
	}

	return result
}

// GeneMutator mutates a gene with the given set of allele
// mutators.
type GeneMutator struct {
	AlleleMutators []alleles.Mutator
}

// Mutate a gene into a new gene
func (gm GeneMutator) Mutate(g Gene) Gene {
	result := Gene{
		make([]alleles.Allele, len(gm.AlleleMutators)),
	}

	for i, m := range gm.AlleleMutators {
		result.Alleles[i] = m.Mutate(g.Alleles[i])
	}

	return result
}
