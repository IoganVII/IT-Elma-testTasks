package arrayOperationService

import "sort"

func arrayRotation(inputArray []int, countRotation int) []int {
	lengthInputArray := len(inputArray)
	for rotationIncrement := 0; rotationIncrement < countRotation; rotationIncrement++ {
		lastElement := inputArray[lengthInputArray-1]

		for i := lengthInputArray - 2; i >= 0; i-- {
			inputArray[i+1] = inputArray[i]
		}

		inputArray[0] = lastElement
	}
	return inputArray
}

func arrayCheckSequence(inputArray []int) int {

	sort.Slice(inputArray[:], func(i, j int) bool {
		return inputArray[i] < inputArray[j]
	})

	for i := 0; i < len(inputArray)-1; i++ {
		if inputArray[i]+1 != inputArray[i+1] {
			return 0
		}
	}
	return 1
}

func arrayFindLoner(inputArray []int) int {

	var hashMap = make(map[int]bool)

	for _, data := range inputArray {
		hashMap[data] = !hashMap[data]
	}

	for key, data := range hashMap {
		if data {
			return key
		}
	}

	return 0
}

func arrayFindSkipEelement(inputArray []int) int {

	sort.Slice(inputArray[:], func(i, j int) bool {
		return inputArray[i] < inputArray[j]
	})

	for i := 0; i < len(inputArray)-1; i++ {
		if inputArray[i]+1 != inputArray[i+1] {
			return inputArray[i] + 1
		}
	}

	return -1

}
