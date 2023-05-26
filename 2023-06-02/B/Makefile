clean:
	rm -rf B B.out

B: B.go
	go build -o B B.go
	./B < B.in > B.out
	diff B.out B.res
