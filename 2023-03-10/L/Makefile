clean:
	rm -rf L L.out

L: L.go
	go build -o L L.go
	./L < L.in > L.out
	diff L.out L.res
