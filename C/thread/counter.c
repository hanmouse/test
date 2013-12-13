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
	int i;
	int e;
	int rv_from_thread[NUM_THREADS];
	pthread_t thread[NUM_THREADS];

	for (i = 0; i < NUM_THREADS; i++) {
		e = pthread_create( &thread[i], NULL, &count, (void *) i);
		if (e) {
			fprintf( stderr, "pthread_create() failed: thread_id=%d, E=%d[ %s ]\n", i, e, strerror( e));
			continue;
		}

		printf( "Thread[%d] created\n", i);
	}

	for (i = 0; i < NUM_THREADS; i++) {
		e = pthread_join( thread[i], (void **) &rv_from_thread[i]);
		if (e) {
			fprintf( stderr, "pthread_join() failed: thread_id=%d, E=%d[ %s ]\n", i, e, strerror( e));
			continue;
		}

		printf( "Thread[%d] joined\n", i);
	}

	printf( "Finally number is %d\n", number);

	return 0;
}

