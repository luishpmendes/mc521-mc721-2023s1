clean:
	rm -rf H H.out

H: H.go
	go build -o H H.go
	./H < H.in > H.out
	diff H.out H.res
