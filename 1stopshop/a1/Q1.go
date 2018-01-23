package Q1

// return list of primes less than N
func sieveOfEratosthenes(N int) (n int) {

	var primes []int

	b := make([]bool, N)

	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)

		for k := i * i; k < N; k += i {
			b[k] = true

		}

	}
	n = len(primes)

	return
}
