clean:
	rm -rf E E.out

E: E.go
	go build -o E E.go
	./E < E.in > E.out
	diff E.out E.res
