clean:
	rm -rf F F.out

F: F.go
	go build -o F F.go
	./F < F.in > F.out
	diff F.out F.res
