package arrayOperationService

func ArrayRotation(inputArray []int, countRotation int) []int {
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

func ArrayCheckSequence(inputArray []int) int {
	minElement := inputArray[0]
	maxElement := inputArray[0]
	isSequence := true

	for i := 1; i < len(inputArray); i++ {
		if inputArray[i] < minElement {
			minElement = inputArray[i]
		}
		if inputArray[i] > maxElement {
			maxElement = inputArray[i]
		}
	}
	var countInput int
	for elemtnValue := minElement; elemtnValue <= maxElement; elemtnValue++ {
		countInput = 0
		for j := 0; j < len(inputArray); j++ {
			if inputArray[j] == elemtnValue {
				countInput++
				if countInput > 1 {
					break
				}
			}
		}
		if countInput == 0 || countInput > 1 {
			isSequence = false
			break
		}
	}

	if isSequence {
		return 1
	} else {
		return 0
	}

}

func ArrayFindLoner(inputArray []int) int {
	var result int
	for i := 0; i < len(inputArray); i++ {

		if i == len(inputArray)-1 {
			return inputArray[i]
		}

		if inputArray[i] == 0 {
			continue
		}

		isFind := false
		temp := inputArray[i]
		for j := i + 1; j < len(inputArray); j++ {
			if inputArray[j] == temp {
				isFind = true
				inputArray[j] = 0
			}
		}
		if !isFind {
			result = inputArray[i]
			break
		}
	}
	return result
}

func ArrayFindSkipEelement(inputArray []int) int {
	minElement := inputArray[0]
	maxElement := inputArray[0]
	result := -1

	for i := 1; i < len(inputArray); i++ {
		if inputArray[i] < minElement {
			minElement = inputArray[i]
		}
		if inputArray[i] > maxElement {
			maxElement = inputArray[i]
		}
	}
	var countInput int
	for elemtnValue := minElement; elemtnValue <= maxElement; elemtnValue++ {
		countInput = 0
		for j := 0; j < len(inputArray); j++ {
			if inputArray[j] == elemtnValue {
				countInput++
				break
			}
		}
		if countInput == 0 {
			result = elemtnValue
			break
		}
	}

	return result

}
