package main

func TwoNumberSum(array []int, target int) []int {
	// Write your code here.
    storedValues := map[int]struct{}{}

    for _, currentValue := range array {
        numNeeded := target - currentValue

        if _, ok := storedValues[numNeeded]; ok {
            return []int{numNeeded, currentValue}
        }

        storedValues[currentValue] = struct{}{}
    } 
	return []int{}
}
