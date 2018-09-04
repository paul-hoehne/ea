package ea

// PopulationStatistics represent basic statistical facts
// about the current simulation
type PopulationStatistics struct {
	MinFitness, MaxFitness, AverageFitness, FitnessDeviation float64
}

// Population is a group of individuals
type Population struct {
	Individuals []Individual
	Factory     IndividualFactory
	Size        int
	Statistics  PopulationStatistics
}

// Initialize a new population of random individuals
func (p *Population) Initialize() {

}

// CreateOffspring creates a set of offspring given the parents,
// using the given selector to select parents.
func (p *Population) CreateOffspring(s Selector) []Individual {
	return nil
}

// Compete pits the offspring against the existing population
// for survival.
func (p *Population) Compete(s Selector, offspring []Individual) *Population {
	return nil
}

// UpdateStatistics updates the population's statistics
func (p *Population) UpdateStatistics() {

}
