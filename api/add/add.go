package add

import "fmt"

func Add(x int, y int, z int) int {
	fmt.Println("adding: ", x, y, z)
	return x + y + z

}

func AddAll(numbers []int) int {
	fmt.Println("adding: ", numbers)
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum

}
