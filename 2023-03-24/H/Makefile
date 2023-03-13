CPP=g++
CARGS=-std=c++17 -Wall -Werror -O0 -g3

clean:
	rm -rf H H.out

H: H.cpp
	$(CPP) $(CARGS) H.cpp -o H
	./H < H.in > H.out
	diff H.out H.res
