package daily

func LastRemaining(n int, m int) int {
	f := 0 //f(1, m)
	for i := 2; i<=n; i+=1 { // f(2,m) = (f(1, m) + m) % 2
		f = (f+m)%i  // f(n, m) = (f(n-1, m) + m)  % n
	}
	return f
}

func LastRemainingRecurve(n int, m int) int {
	return recurve(n, m)
}

func recurve(n int, m int) int {
	if n == 1 {
		return 0
	}
	x := recurve(n-1, m)
	return (x + m) % n
}
