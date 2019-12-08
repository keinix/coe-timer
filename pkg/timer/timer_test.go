package timer

import (
	"testing"
)

func TestGenerateEventMap(t *testing.T) {
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
		{name: "3 elements; start: index 0; want: index 2", start: Cold, want: Poison, expected:
			map[int]Event{8: Enter, 12: Exit, 7: Warning, 6: Warning, 5: Warning, 9: Active, 10: Active, 11: Active}},
		{name: "3 elements; start: index 1; want: index 0", start: Physical, want: Cold, expected:
			map[int]Event{8: Enter, 12: Exit, 7: Warning, 6: Warning, 5: Warning, 9: Active, 10: Active, 11: Active}},
		{name: "3 elements; start: index 1; want: index 1", start: Physical, want: Physical, expected:
			map[int]Event{0: Enter, 4: Exit, 11: Warning, 10: Warning, 9: Warning, 1: Active, 2: Active, 3: Active}},
		{name: "3 elements; start: index 1; want: index 2", start: Physical, want: Poison, expected:
			map[int]Event{4: Enter, 8: Exit, 3: Warning, 2: Warning, 1: Warning, 5: Active, 6: Active, 7: Active}},
		{name: "3 elements; start: index 2; want: index 0", start: Poison, want: Cold, expected:
			map[int]Event{4: Enter, 8: Exit, 3: Warning, 2: Warning, 1: Warning, 5: Active, 6: Active, 7: Active}},
		{name: "3 elements; start: index 2; want: index 1", start: Poison, want: Physical, expected:
			map[int]Event{8: Enter, 12: Exit, 7: Warning, 6: Warning, 5: Warning, 9: Active, 10: Active, 11: Active}},
		{name: "3 elements; start: index 2; want: index 2", start: Poison, want: Poison, expected:
			map[int]Event{0: Enter, 4: Exit, 11: Warning, 10: Warning, 9: Warning, 1: Active, 2: Active, 3: Active}},
	}

	for _, test := range tables {
		actual, err := generateEventMap(test.start, test.want, cycle)
		if err != nil {
			t.Errorf("error in test: %v", err)
			t.FailNow()
		}
		compareEventMap(t, test.name, actual, test.expected)
	}
}

func compareEventMap(t *testing.T, caseName string, actual map[int]Event, expected map[int]Event) {
	if len(actual) != len(expected) {
		t.Errorf("case: %v; actual size (%d) does not match expected (%d) \n expected: %v \n actual %v",
			caseName, len(actual), len(expected),
			convertToHumanReadableOutput(expected), convertToHumanReadableOutput(actual))
		t.FailNow()
	}
	for time, expectedVal := range expected {
		actualVal, ok := actual[time]
		if !ok || actualVal != expectedVal {
			t.Errorf("case: %v; Expected does not match actual \n expected: %v \n actual: %v",
				caseName, convertToHumanReadableOutput(expected), convertToHumanReadableOutput(actual))
		}
	}
}

func convertToHumanReadableOutput(eventMap map[int]Event) map[int]string {
	var output = make(map[int]string)
	for time, event := range eventMap {
		output[time] = getEventName(event)
	}
	return output
}

func getEventName(e Event) string {
	switch e {
	case Enter:
		return "Enter"
	case Exit:
		return "Exit"
	case Active:
		return "Active"
	case Warning:
		return "Warning"
	}
	return "Event not found"
}
