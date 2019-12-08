package timer

import (
	"errors"
	"fmt"
)

type Event uint8
type Element uint8
type Class uint8

const (
	Warning Event = iota
	Enter
	Active
	Exit
)

const(
	Arcane Element = iota
	Cold
	Fire
	Holy
	Lightning
	Physical
	Poison
)

const(
	Barbarian Class = iota
	Crusader
	DemonHunter
	Monk
	Necromancer
	WitchDoctor
	Wizard
)

const eleDuration int = 4 // seconds

func Start(c chan<- Event, class Class, startEle Element, wantEle Element) error {
	cycle, err := getCycle(class)
	if err != nil {
		return err
	}
	cycleLength := len(cycle) * eleDuration
	eventMap, err := generateEventMap(startEle, wantEle, cycle)
	if err != nil {
		return err
	}
	if event, ok := eventMap[0]; ok {
		updateEventTime(&eventMap, 0, cycleLength)
		c <- event
	}
	//t := time.NewTicker(1e9) // tick every second
	return nil
}

func getCycle(c Class) ([]Element, error) {
	switch c {
	case Barbarian:
		return []Element{Cold, Fire, Lightning, Physical}, nil
	case Crusader:
		return []Element{Fire, Holy, Lightning, Physical}, nil
	case DemonHunter:
		return []Element{Cold, Fire, Lightning, Physical}, nil
	case Monk:
		return []Element{Cold, Fire, Holy, Lightning, Physical}, nil
	case Necromancer:
		return []Element{Cold, Physical, Poison}, nil
	case WitchDoctor:
		return []Element{Cold, Fire, Physical, Poison}, nil
	case Wizard:
		return []Element{Arcane, Cold, Fire, Lightning}, nil
	}
	return nil, fmt.Errorf("%v is not a valid class", c)
}

// Events mapped to the number of seconds from the ticker's start they should happen
func generateEventMap(startEle Element, wantEle Element, cycle []Element) (map[int]Event, error) {
	enter, err := getEnterTime(startEle, wantEle, cycle)
	if err != nil {
		return nil, err
	}
	exit := enter + eleDuration // exit is the start of the next Element
	warningTics := getWarningTics(enter, cycle)
	activeTics := getActiveTics(enter)

	eventMap := make(map[int]Event)
	eventMap[enter] = Enter
	eventMap[exit] = Exit
	for _, tic := range activeTics {
		eventMap[tic] = Active
	}
	for _, tic := range warningTics {
		eventMap[tic] = Warning
	}
	return eventMap, nil
}

func getActiveTics(start int) []int {
	activeTics := make([]int, 3)
	for i := 0; i < len(activeTics); i++ {
		activeTics[i] = start + i + 1
	}
	return activeTics
}

func getWarningTics(start int, cycle []Element) []int {
	cycleLength := len(cycle) * eleDuration
	warningTics := make([]int, 3)
	for i := 0; i < len(warningTics); i++ {
		tic := start - i - 1
		if tic < 0 {
			// time for warning tic has already passed for this rotation and should
			// be played for the first time next rotation
			tic += cycleLength
		}
		warningTics[i] = start
	}
	return warningTics
}

// Modify the Event time for the next rotation
func updateEventTime(m *map[int]Event, time int, cycleLength int) {
	e := (*m)[time]
	delete(*m, time)
	(*m)[time + cycleLength] = e
}

// The time in seconds from the startEle to wantEle
func getEnterTime(startEle Element, wantEle Element, cycle []Element) (int, error) {
	startIndex, err := getElementIndex(cycle, startEle)
	if err != nil {
		return 0, err
	}
	wantIndex, err := getElementIndex(cycle, wantEle)
	if err != nil {
		return 0, err
	}
	elementsUntilWantEle := wantIndex - startIndex
	var enter int
	if elementsUntilWantEle >= 0 {
		// the wanted element is after the start element in the current rotation
		enter = elementsUntilWantEle * eleDuration
	} else {
		// the wanted element won't be entered until the next rotation
		enter = (elementsUntilWantEle + len(cycle)) * eleDuration
	}
	return enter, nil
}

func getElementIndex(cycle []Element, e Element) (int, error) {
	for i, val := range cycle {
		if val == e {
			return i, nil
		}
	}
	return -1, errors.New("element not in cycle")
}