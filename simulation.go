package ea

import (
	"sync"
)

// Stopper returns true when it's time to stop a simulation
type Stopper interface {
	Stop(s *Simulation) bool
}

// MaxGenerationsStopper indicates it's time to stop the simulation
// after a given number of generations.
type MaxGenerationsStopper struct {
	Max int
}

// Stop indicates that the simulation should stop after it
// exceedes the given number of generations.
func (m MaxGenerationsStopper) Stop(s *Simulation) bool {
	return false
}

// Simulation holds the basic parameters of the simulation
type Simulation struct {
	ParentSelector      Selector
	CompetitionSelector Selector
	Generations         int
	Population          *Population
	StopCriteria        []Stopper
	FitnessFunction     FitnessFunction
}

// Start initializes the starting population
func (s *Simulation) Start() {

}

// Generation executes one generational epoch
func (s *Simulation) Generation() {
	offspring := s.Population.CreateOffspring(s.ParentSelector)

	var wg sync.WaitGroup
	for _, o := range offspring {
		go s.FitnessFunction(&o, &wg)
	}
	wg.Wait()

	s.Population = s.Population.Compete(s.CompetitionSelector, offspring)
	s.Population.UpdateStatistics()
}

// Run the simulation until the stop criteria is met
func (s *Simulation) Run() {
	for {
		s.Generation()

		for _, st := range s.StopCriteria {
			if st.Stop(s) {
				break
			}
		}
	}
}
