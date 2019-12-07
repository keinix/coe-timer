package timer

import "testing"

func TestGetEventMap(t *testing.T) {
	cycle := []Element{Cold, Physical, Poison}
	tables := []struct {
		name string
		start Element
		want Element
		expected map[int]Event
	} {
		{name: "3 elements; start: index 0; want: index 0", start: Cold, want: Cold, expected:
			map[int]Event{0: Enter, 4: Exit, 11: Warning, 10: Warning, 9: Warning, 1: Active, 2: Active, 3: Active}},
		{name: "3 elements; start: index 0; want: index 1", start: Cold, want: Physical, expected:
			map[int]Event{4: Enter, 8: Exit, 3: Warning, 2: Warning, 1: Warning, 5: Active, 6: Active, 7: Active}},


	}

	//{name: "3 elements; start: index 0; want: index 2", start: Cold, want: Poison},
	//{name: "3 elements; start: index 1; want: index 0", start: Physical, want: Cold},
	//{name: "3 elements; start: index 1; want: index 1", start: Physical, want: Physical},
	//{name: "3 elements; start: index 1; want: index 2", start: Physical, want: Poison},
	//{name: "3 elements; start: index 2; want: index 0", start: Poison, want: Cold},
	//{name: "3 elements; start: index 2; want: index 1", start: Poison, want: Physical},
	//{name: "3 elements; start: index 2; want: index 2", start: Poison, want: Poison},
}
