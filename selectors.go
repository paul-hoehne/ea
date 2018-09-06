package ea

import "math/rand"

// Selector is a way to choose a set of individuals
type Selector interface {
	Select(pop Population) []Individual
}

// FitnessProportionalSelector selects one or more individuals
// based on fitness proportionality.
type FitnessProportionalSelector struct {
	WithReplacement bool
	SelectCount     int
}

func (fp FitnessProportionalSelector) selectWithReplacement(pop []Individual) []Individual {
	result := make([]Individual, 0, fp.SelectCount)

	sumFitness := 0.0
	for _, i := range pop {
		sumFitness += i.Fitness
	}

	for j := 0; j < fp.SelectCount; j++ {
		atRandom := rand.Float64() * sumFitness
		for _, i := range pop {
			if atRandom <= i.Fitness {
				result = append(result, i)
				break
			} else {
				atRandom -= i.Fitness
			}
			if atRandom < 0.0 || len(result) == fp.SelectCount {
				break
			}
		}
	}

	return result
}

func (fp FitnessProportionalSelector) selectWithoutReplacement(pop []Individual) []Individual {
	result := make([]Individual, 0, fp.SelectCount)

	sumFitness := 0.0
	for _, i := range pop {
		sumFitness += i.Fitness
	}

	indexes := make(map[int]int, len(pop))
	for i := range pop {
		indexes[i] = i
	}

	for j := 0; j < fp.SelectCount; j++ {
		atRandom := rand.Float64() * sumFitness
		for idx := range indexes {
			if atRandom <= pop[idx].Fitness {
				result = append(result, pop[idx])
				sumFitness -= pop[idx].Fitness
				delete(indexes, idx)
				break
			} else {
				atRandom -= pop[idx].Fitness
			}
			if atRandom < 0.0 || len(result) == fp.SelectCount {
				break
			}
		}
	}

	return result
}

// Select individuals based on their fitness.
func (fp FitnessProportionalSelector) Select(pop []Individual) []Individual {
	if fp.WithReplacement {
		return fp.selectWithReplacement(pop)
	}
	return fp.selectWithoutReplacement(pop)
}

// UniformSelector selects individuals on an eqiprobable basis.
type UniformSelector struct {
	WithReplacement bool
	SelectCount     int
}

// Select individuals using a uniform distribution
func (us UniformSelector) Select(pop []Individual) []Individual {
	return nil
}
