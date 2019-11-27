package fibonacci

func Fibonacci(n uint64) uint64{
	if n == 0 {
		return 0
	}

	if n <= 2 {
		return 1
	}

	return Fibonacci(n - 1) + Fibonacci(n - 2)
}

