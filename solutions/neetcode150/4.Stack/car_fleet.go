package main

import "sort"

type Car struct {
	Position int
	Time     float64
}

// Leetcode #853 
func carFleet(target int, position []int, speed []int) int {
	var cars []Car

	for i := 0; i < len(position); i++ {
		time := float64(target-position[i]) / float64(speed[i])
		cars = append(cars, Car{Position: position[i], Time: time})
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Position > cars[j].Position
	})

	var stack []float64

	for _, car := range cars {
		// If the current car takes longer or the same time to reach the target than the last car in the stack,
		// it will not form a new fleet but join the fleet of the car in the stack.
		if len(stack) > 0 && car.Time <= stack[len(stack)-1] {
			continue
		}
		// Otherwise, this car leads a new fleet.
		stack = append(stack, car.Time)
	}

	return len(stack) // The number of fleets is the size of the stack
}