clean:
	rm -rf K K.out

K: K.go
	go build -o K K.go
	./K < K.in > K.out
	diff K.out K.res
