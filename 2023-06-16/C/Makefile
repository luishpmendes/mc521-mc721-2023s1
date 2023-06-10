clean:
	rm -rf C C.out

C: C.go
	go build -o C C.go
	./C < C.in > C.out
	diff C.out C.res
