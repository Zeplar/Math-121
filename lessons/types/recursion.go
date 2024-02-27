package main

func fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func sumto(n int) int {
	return n + sumto(n-1)
}

func main() {
	for i := 1; i < 10; i++ {
		i := i
		println(fib(i))
	}

	// println(sumto(10))

}

/*
	Partitions: How many ways to write n as a sum?

5 = 5
5 = 4 + 1
5 = 3 + 2
5 = 3 + 1 + 1
5 = 2 + 2 + 1
5 = 2 + 1 + 1 + 1
5 = 1 + 1 + 1 + 1 + 1
*/

// Partitions of n using numbers up to m
func part_under(n, m int) int

// 					 partitions not using m  + partitions using m
// part_under(n,m) == part_under(n, m-1) + part_under(n-m, m)

func partitition(n int) int {
	return part_under(n, n)
}
