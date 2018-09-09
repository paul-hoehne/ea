package ea

import (
	"math/rand"
	"sync"

	"github.com/chilts/sid"

	"github.com/paul-hoehne/ea/alleles"
	"github.com/paul-hoehne/ea/genes"
)

// FitnessFunction is a function that takes and individual,
// calculates their fitness, and indicates it's done to the wait
// group.  For non-trivial problems like neural network optimization
// this may take a long time to compute (and therefore) it is
// expected to run in separate goroutine.
type FitnessFunction func(i *Individual, wg *sync.WaitGroup)

// ReproductionStrategy takes a set of individuals and produces
// a new individual
type ReproductionStrategy func([]Individual) []Individual

// Individual is a member of the population
type Individual struct {
	ID      string
	Genes   []genes.Gene
	Fitness float64
}

// IndividualFactory creates a new individual
type IndividualFactory struct {
	GeneFactories []genes.GeneFactory
}

// IndividualMutators contains the gene mutators necessary to
// perform mutation on an individual.
type IndividualMutators struct {
	GeneMutators []genes.GeneMutator
}

// Create a new individual
func (i IndividualFactory) Create() Individual {
	result := Individual{
		ID:    sid.Id(),
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
		ID:    sid.Id(),
		Genes: make([]genes.Gene, len(i.Genes)),
	}

	for idx, g := range i.Genes {
		result.Genes[idx] = g.Copy()
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
		ID:    sid.Id(),
		Genes: make([]genes.Gene, len(i.Genes)),
	}

	for idx := range i.Genes {
		if rand.Float64() < crossoverRate {
			result.Genes[idx] = other.Genes[idx].Copy()
		} else {
			result.Genes[idx] = i.Genes[idx].Copy()
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

// SimpleCrossoverReproductionStrategy creates a func that reproduces
// using crossover and the given crossover rate.
func SimpleCrossoverReproductionStrategy(crossover float64, numOffspring int) ReproductionStrategy {
	return func(matingPair []Individual) []Individual {
		result := make([]Individual, numOffspring)

		for j := 0; j < numOffspring; j++ {
			i := Individual{
				ID:    sid.Id(),
				Genes: make([]genes.Gene, len(matingPair[0].Genes)),
			}
			for g := range matingPair[0].Genes {
				if rand.Float64() < crossover {
					i.Genes[g] = matingPair[1].Genes[g]
				} else {
					i.Genes[g] = matingPair[0].Genes[g]
				}
			}
			result[j] = i
		}

		return result
	}
}

// SimpleCopyReproductionStrategy Takes a set of individual and
// (round-robin fashion) returns copies of thos individuals.
// (Used mainly for asexual reproduction in something like a
// classic Evolutionary Strategy)
func SimpleCopyReproductionStrategy(numOffspring int) ReproductionStrategy {
	return func(individual []Individual) []Individual {
		result := make([]Individual, numOffspring)

		for i := range result {
			result[i].ID = sid.Id()

			result[i].Genes = make([]genes.Gene, len(individual[0].Genes))
			for gidx, g := range individual[0].Genes {
				result[i].Genes[gidx] = g.Copy()
			}
		}

		return result
	}
}
