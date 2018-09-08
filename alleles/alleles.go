package alleles

// Allele is a value expression of a gene.
type Allele interface {
	Copy() Allele
}

// Factory produces new random alleles.
type Factory interface {
	Random() Allele
}

// Mutator is responsible for create a new
// instance of a mutator function
type Mutator interface {
	Mutate(Allele) Allele
}
