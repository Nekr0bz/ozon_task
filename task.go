package ozon_task

func max(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

func maxOnesAfterRemoveItem(bytes []byte) uint {
	var result, prevSum, currentSum, zeroFlag uint = 0, 0, 0, 1

	for i, b := range bytes{
		if b == 1 {
			currentSum += 1
			if i == len(bytes)-1 {
				result = max(result, currentSum+prevSum)
			}
		} else {
			zeroFlag = 0
			if i != 0 && bytes[i-1] == 0 {
				prevSum = 0
			} else {
				result = max(result, currentSum+prevSum)
				prevSum = currentSum
				currentSum = 0
			}
		}
	}
	return result - zeroFlag
}


