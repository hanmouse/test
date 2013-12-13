CC = gcc
libs = -lpthread
CFLAGS = -g

all: counter mutex_attr condvar thread_local_data

counter_objs = counter.o

counter : $(counter_objs)
	$(CC) -o counter $(counter_objs) $(libs)

counter.o : counter.c
	$(CC) -c counter.c

mutex_attr_objs = mutex_attr.o

mutex_attr : $(mutex_attr_objs)
	$(CC) -o mutex_attr $(mutex_attr_objs) $(libs)

mutex_attr.o : mutex_attr.c
	$(CC) -c mutex_attr.c

condvar_objs = condvar.o

condvar : $(condvar_objs)
	$(CC) -o condvar $(condvar_objs) $(libs)

condvar.o : condvar.c
	$(CC) -c condvar.c

thread_local_data_objs = thread_local_data.o

thread_local_data : $(thread_local_data_objs)
	$(CC) -o thread_local_data $(thread_local_data_objs) $(libs)

thread_local_data.o : thread_local_data.c
	$(CC) -c thread_local_data.c

clean :
	rm *.o counter mutex_attr condvar thread_local_data
