package add

import "fmt"

func Add(x int, y int, z int) int {
	fmt.Println("adding: ", x, y, z)

	return x + y + z
}

func AddAll(numbers []int) int {
	fmt.Println("AddAll, adding: ", numbers)

	var sum int
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumPositive(numbers []int) int {
	fmt.Println("SumPositive, adding: ", numbers)

	var sum int
	for _, number := range numbers {
		if number < 0 {
			fmt.Println("SumPositive, -ve number found: ", number)
			sum = -1
			break
		}

		sum += number
	}

	return sum
}
