package main

import (
	"fmt"
	"math/rand"
)

// This program simulates a game show where a contestant is given the option to pick one of three doors.
// Two doors contain goats, and the other contains a car. The game show host knows what is behind each door
// ahead of time. After the contestant picks a door, the game show host reveals a goat by opening one of the
// doors that the contestant did NOT pick. The game show host then gives the contestant the option to switch
// which door they picked. Since there are two doors remaining, many people think there is a 50% chance for
// each door to contain the car, but in reality, there is actually a 33% chance for the car to be behind the door that the
// contestant originally picked and a 66% chance for the car to be behind the other door. This program runs 10,000 trials
// and prints out the results.

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

	// Out of the two doors that the contestant did NOT choose, the game show host will open one of them
	// that has a goat behind it. If they both have goats, then it does not matter which one the
	// host opens. The contestant is now given the option to switch which door they picked.
	// As can be observed with this experiment, the contestant should choose to switch doors because 66% of
	// the time, the goat is behind the door that they did not originally pick.
	//
	// It's also easier to understand why this is the case once you see this situation simulated with code.
	// I didn't even need to write any code to simulate a game show host opening a door because it has no
	// impact on whether the car was behind the door that the contestant originally picked.
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
