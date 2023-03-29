clean:
	rm -rf G G.out

G: G.go
	go build -o G G.go
	./G < G.in > G.out
	diff G.out G.res
