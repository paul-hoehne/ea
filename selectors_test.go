package ea

import (
	"testing"

	"github.com/chilts/sid"
)

func TestFitnessProportionalSelector(t *testing.T) {
	s := FitnessProportionalSelector{
		SelectCount: 1,
	}

	pop := []Individual{
		Individual{Fitness: 1.0},
		Individual{Fitness: 0.0},
	}

	selected := s.Select(pop)

	if len(selected) != 1 {
		t.Errorf("Expected 1 individual but got %d", len(selected))
	}

	if selected[0].Fitness != 1.0 {
		t.Errorf("Expected fitness to be 1.0 but got %f", selected[0].Fitness)
	}
}

func TestFitnessProportionalSelectorWithoutReplacement(t *testing.T) {
	s := FitnessProportionalSelector{
		SelectCount:     10,
		WithReplacement: false,
	}

	pop := []Individual{
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
	}

	foundIds := make(map[string]string, 10)

	for _, i := range pop {
		if i.Fitness == 1.0 {
			foundIds[i.ID] = i.ID
		}
	}

	selected := s.Select(pop)

	if len(selected) != 10 {
		t.Errorf("Expected 10 selected individuals but got %d", len(selected))
	}

	for _, i := range selected {
		if i.Fitness < 1.0 {
			t.Error("Individual selected did not have adequate fitness")
		}

		if _, ok := foundIds[i.ID]; !ok {
			t.Errorf("Unable to find id in the list of selected IDs")
		}

		delete(foundIds, i.ID)
	}

	if len(foundIds) > 0 {
		t.Errorf("Expected foundIds to be empty but got %d", len(foundIds))
	}
}

func TestFitnessProportionalSelectorWithReplacement(t *testing.T) {
	s := FitnessProportionalSelector{
		SelectCount:     10,
		WithReplacement: true,
	}

	pop := []Individual{
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 0.0},
	}

	selected := s.Select(pop)

	if len(selected) > 10 {
		t.Errorf("Expected 10 or fewer individuals but got %d", len(selected))
	}

	for _, i := range selected {
		if i.Fitness < 1.0 {
			t.Error("Individual selected did not have adequate fitness")
		}
	}
}

func TestUniformSelectorWithReplacement(t *testing.T) {
	s := UniformSelector{
		SelectCount:     5,
		WithReplacement: true,
	}

	pop := []Individual{
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
	}

	selected := s.Select(pop)

	if len(selected) != s.SelectCount {
		t.Errorf("Expcted %d but got %d", s.SelectCount, len(selected))
	}
}

func TestUniformSelectorWithoutReplacement(t *testing.T) {
	s := UniformSelector{
		SelectCount: 5,
	}

	pop := []Individual{
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
		Individual{ID: sid.Id(), Fitness: 1.0},
	}

	selected := s.Select(pop)

	if len(selected) != len(pop) {
		t.Errorf("Expcted %d but got %d", len(pop), len(selected))
	}
}
