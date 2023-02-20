clean:
	rm -rf J J.out

J: J.go
	go build -o J J.go
	./J < J.in > J.out
	diff J.out J.res
