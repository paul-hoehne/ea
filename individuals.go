package ea

import (
	"math/rand"
	"sync"

	"github.com/paul-hoehne/ea/alleles"
	"github.com/paul-hoehne/ea/genes"
)

// FitnessFunction is a function that takes and individual,
// calculates their fitness, and indicates it's done to the wait
// group.  For non-trivial problems like neural network optimization
// this may take a long time to compute (and therefore) it is
// expected to run in separate goroutine.
type FitnessFunction func(i *Individual, wg *sync.WaitGroup)

// Individual is a member of the population
type Individual struct {
	Genes   []genes.Gene
	Fitness float64
}

// IndividualFactory creates a new individual
type IndividualFactory struct {
	GeneFactories []genes.GeneFactory
}

// Create a new individual
func (i IndividualFactory) Create() Individual {
	result := Individual{
		Genes: make([]genes.Gene, len(i.GeneFactories)),
	}

	for i, f := range i.GeneFactories {
		result.Genes[i] = genes.Gene{
			Alleles: make([]alleles.Allele, len(f.AlleleFactories)),
		}

		for j, af := range f.AlleleFactories {
			result.Genes[i].Alleles[j] = af.Random()
		}
	}

	return result
}

// Spawn is asexual reproduction.
func (i Individual) Spawn() Individual {
	result := Individual{
		Genes: make([]genes.Gene, len(i.Genes)),
	}

	for idx, g := range i.Genes {
		result.Genes[idx] = genes.Gene{
			Alleles: make([]alleles.Allele, len(g.Alleles)),
		}

		for j, a := range g.Alleles {
			result.Genes[idx].Alleles[j] = a
		}
	}

	return result
}

// Breed is sexual reproduction.  Note that crossover is done
// at the gene level.  Since our genes can be code values or real
// numbers as well as bit strings, bit by bit crossover is ill
// defined for this implementation (and may be handled as an optimized
// bitstring individual later)
func (i Individual) Breed(other Individual, crossoverRate float64) Individual {
	result := Individual{
		Genes: make([]genes.Gene, len(i.Genes)),
	}

	for idx, g := range i.Genes {
		result.Genes[idx].Alleles = make([]alleles.Allele, len(g.Alleles))
		for j, a := range g.Alleles {
			if rand.Float64() < crossoverRate {
				result.Genes[idx].Alleles[j] = other.Genes[idx].Alleles[j]
			} else {
				result.Genes[idx].Alleles[j] = a
			}
		}
	}

	return result
}

// Mutate mutates an individual
func (i *Individual) Mutate(muts []genes.GeneMutator, mutationRate float64) {
	for idx, g := range i.Genes {

		if rand.Float64() < mutationRate {
			i.Genes[idx] = muts[idx].Mutate(g)
		}
	}
}
