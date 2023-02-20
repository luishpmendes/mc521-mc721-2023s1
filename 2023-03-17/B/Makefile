CPP=g++
CARGS=-std=c++17 -Wall -Werror -O0 -g3

clean:
	rm -rf B B.out

B: B.cpp
	$(CPP) $(CARGS) B.cpp -o B
	./B < B.in > B.out
	diff B.out B.res
