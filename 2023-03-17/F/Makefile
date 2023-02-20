CPP=g++
CARGS=-std=c++17 -Wall -Werror -O0 -g3

clean:
	rm -rf F F.out

F: F.cpp
	$(CPP) $(CARGS) F.cpp -o F
	./F < F.in > F.out
	diff F.out F.res
