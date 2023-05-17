clean:
	rm -rf I I.out

I: I.go
	go build -o I I.go
	./I < I.in > I.out
	diff I.out I.res
