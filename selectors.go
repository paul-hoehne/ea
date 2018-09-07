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

func (us UniformSelector) selectWithReplacement(pop []Individual) []Individual {
	result := make([]Individual, 0)

	for i := 0; i < us.SelectCount; i++ {
		idx := rand.Int() % len(pop)
		result = append(result, pop[idx])
	}

	return result
}

func (us UniformSelector) selectWithoutReplacement(pop []Individual) []Individual {
	altmap := make([]*Individual, len(pop))
	for i := range pop {
		altmap[i] = &pop[i]
	}

	result := []Individual{}
	for j := 0; j < us.SelectCount && len(altmap) > 0; j++ {
		idx := rand.Int() % len(altmap)

		result = append(result, *(altmap[idx]))
		if idx == 0 {
			altmap = altmap[1:]
		} else if idx == len(altmap)-1 {
			altmap = altmap[:len(altmap)-1]
		} else {
			altmap = append(altmap[:idx], altmap[idx+1:]...)
		}
	}
	return result
}

// Select individuals using a uniform distribution
func (us UniformSelector) Select(pop []Individual) []Individual {
	if us.WithReplacement {
		return us.selectWithReplacement(pop)
	}

	return us.selectWithoutReplacement(pop)
}
