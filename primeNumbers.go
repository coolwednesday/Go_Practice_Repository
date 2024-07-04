package main

// isprime function checks the number is prime or not
func isprime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true //if both failed then num is prime
}

// PrimeNumbers function take an integer and return the count of prime numbers
func PrimeNumbers(n int) int {
	count := 0
	for i := 0; i <= n; i++ {
		if isprime(i) == true {
			count++
		}
	}
	return count

}
