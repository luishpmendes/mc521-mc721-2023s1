clean:
	rm -rf M M.out

M: M.go
	go build -o M M.go
	./M < M.in > M.out
	diff M.out M.res
