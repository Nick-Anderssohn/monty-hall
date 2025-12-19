package main

import (
	"fmt"
	"math/rand"
)

type trialResult struct {
	carIsBehindDoorOriginallyPicked bool
}

type doorResult int

const (
	goat doorResult = iota
	car
)

func main() {
	runAllTrials()
}

func runAllTrials() {
	var numTimesBehindOriginalDoor, numTimesBehindOtherDoor int
	numTrials := 10_000

	for i := 0; i < numTrials; i++ {
		result := runTrial()

		if result.carIsBehindDoorOriginallyPicked {
			numTimesBehindOriginalDoor++
		} else {
			numTimesBehindOtherDoor++
		}
	}

	fmt.Printf("Number of trials: %d\n", numTrials)
	fmt.Printf("Number of times the car was behind the door originally picked: %d\n", numTimesBehindOriginalDoor)
	fmt.Printf("Number of times the car was behind the other door: %d\n", numTimesBehindOtherDoor)
	fmt.Printf("Probability of the car being behind the door originally picked: %.2f%%\n", float32(numTimesBehindOriginalDoor)/float32(numTrials)*100)
}

func runTrial() (result trialResult) {
	doors := generateDoors()

	// contestant chooses a random door
	doorChosen := rand.Intn(3)

	// Out of the two doors that the contestant did NOT choose, the host will open one of them
	// that has a goat behind it (if they both have goats, then it does not matter which one the
	// host opens). The contestant is now given the option to switch which door they picked.
	// As can be observed with this experiment, the contestant should switch doors because 66% of
	// the time, the goat is behind the other door. It's also easier to understand why this is the
	// case once you see this situation simulated with code.
	if doors[doorChosen] == car {
		result.carIsBehindDoorOriginallyPicked = true
	} else {
		result.carIsBehindDoorOriginallyPicked = false
	}

	return result
}

// generateDoors generates three doors. Two contain goats, and one contains a car.
func generateDoors() (doors [3]doorResult) {
	carIndex := rand.Intn(3)
	doors[carIndex] = car
	return doors
}
