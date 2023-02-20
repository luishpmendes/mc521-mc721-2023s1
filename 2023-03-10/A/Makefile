clean:
	rm -rf A A.out

A: A.go
	go build -o A A.go
	./A < A.in > A.out
	diff A.out A.res
