clean:
	rm -rf D D.out

D: D.go
	go build -o D D.go
	./D < D.in > D.out
	diff D.out D.res
