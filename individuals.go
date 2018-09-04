package ea

import (
	"sync"

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

// Spawn is asexual reproduction.
func (i Individual) Spawn() Individual {
	return Individual{}
}

// Breed is sexual reproduction
func (i Individual) Breed(crossoverRate float64) Individual {
	return Individual{}
}

// Mutate mutates an individual
func (i *Individual) Mutate(mutationRate float64) {

}
