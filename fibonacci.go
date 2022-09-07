package fibonacci

func Loop(num int) int {
	if num < 2 {
		return num
	}

	num1 := 0
	num2 := 1

	for i := 2; i <= num; i++ {
		num2, num1 = num2+num1, num2
	}

	return num2
}

func Recursive(num int) int {
	if num < 2 {
		return num
	}

	return Recursive(num-1) + Recursive(num-2)
}

func RecursiveCache(num int) int {
	seen := map[int]int{
		0: 0,
		1: 1,
	}

	var recursiveFibonacci func(int) int

	recursiveFibonacci = func(num int) int {
		_, ok := seen[num]
		if !ok {
			seen[num] = recursiveFibonacci(num-1) + recursiveFibonacci(num-2)
		}

		return seen[num]
	}

	return recursiveFibonacci(num)
}
