#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <errno.h>
#include <time.h>
#include <stdbool.h>
#include <pthread.h>
#include <sys/time.h>

#define NUM_THREADS	10

int number = 0;
pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;

void *count( void *param)
{
	int i;
	int id = (int) param;
	int e;
	struct timeval tv;
	useconds_t usleep_time;

	for (i = 0; i < 10; i++) {

		e = gettimeofday( &tv, NULL);
		if (e < 0) {
			fprintf( stderr, "gettimeofday() failed: E=%d[ %s ]\n", e, strerror( e));
			continue;
		}

		e = pthread_mutex_lock( &mutex);
		if (e) {
			fprintf( stderr, "pthread_mutex_lock() failed: E=%d[ %s ]\n", e, strerror( e));
			break;
		}

		number++;

		printf( "[thread %d] number=%d\n", id, number);
		usleep_time = rand_r( &tv.tv_usec) % 10000;
		usleep( usleep_time);

		e = pthread_mutex_unlock( &mutex);
		if (e) {
			fprintf( stderr, "pthread_mutex_unlock() failed: E=%d[ %s ]\n", e, strerror( e));
			break;
		}
	}


	return (void *) id;
}

int main( int argc, char **argv)
{
	int i = 0;
	int e;
	int rv_from_thread;
	pthread_t thread;
	pthread_mutex_t mutex;
	pthread_mutexattr_t mutex_attr;

	e = pthread_mutexattr_init( &mutex_attr);
	if (e) {
		fprintf( stderr, "pthread_mutexattr_init() failed: E=%d[ %s ]\n", e, strerror( e));
		return 1;
	}

	pthread_mutexattr_setpshared( &mutex_attr, PTHREAD_PROCESS_SHARED);

	e = pthread_create( &thread, NULL, &count, (void *) i);
	if (e) {
		fprintf( stderr, "pthread_create() failed: E=%d[ %s ]\n", e, strerror( e));
		return 1;
	}

	printf( "Thread[%d] created\n", i);

	e = pthread_join( thread, (void **) &rv_from_thread);
	if (e) {
		fprintf( stderr, "pthread_join() failed: E=%d[ %s ]\n", e, strerror( e));
		return 1;
	}

	return 0;
}

