CPP=g++
CARGS=-std=c++17 -Wall -Werror -O0 -g3

clean:
	rm -rf A A.out

A: A.cpp
	$(CPP) $(CARGS) A.cpp -o A
	./A < A.in > A.out
	diff A.out A.res
