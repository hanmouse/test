CC = gcc
libs = -lpthread
CFLAGS = -g

all: cirqueue_test

objs = main.o reader.o queue.o file.o sighandle.o thread_life.o

cirqueue_test : $(objs)
	$(CC) -o cirqueue_test $(objs) $(libs)

main.o : main.c
	$(CC) -c main.c

reader.o : reader.c
	$(CC) -c reader.c

queue.o : queue.c
	$(CC) -c queue.c

file.o : file.c
	$(CC) -c file.c

sighandle.o : sighandle.c
	$(CC) -c sighandle.c

thread_life.o : thread_life.c
	$(CC) -c thread_life.c

clean :
	rm *.o cirqueue_test
