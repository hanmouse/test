#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <string.h>
#include <strings.h>
#include <time.h>
#include <errno.h>
#include <stdarg.h>
#include <stdbool.h>
#include <ctype.h>
#include <signal.h>
#include <pthread.h>

#define NUM_THREADS	10

__thread int count;
//int count;

void *count_func( void *param)
{
	int thread_id = (int) param;
	int i;

	for (i = 0; i < 10; i++) {
		count++;
		fprintf( stderr, "[thread %d] count=%d\n", thread_id, count);
		usleep( 1000);
	}

	return NULL;
}

int main( int argc, char **argv)
{
	int i;
	pthread_t threads[NUM_THREADS];

	for (i = 0; i < NUM_THREADS; i++) {
		pthread_create( &threads[i], NULL, count_func, (void *) i);
	}

	for (i = 0; i < NUM_THREADS; i++) {
		pthread_join( &threads[i], NULL);
	}

	return 0;
}

