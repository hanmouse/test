#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <errno.h>
#include <time.h>
#include <stdbool.h>
#include <pthread.h>
#include <sys/time.h>

#define QUEUE_LENGTH	200

int queue[QUEUE_LENGTH];
int queue_len = 0;
pthread_mutex_t mutex = PTHREAD_MUTEX_INITIALIZER;
pthread_cond_t cv;

void *producer( void *param)
{
	int i;

	for (i = 0; i < QUEUE_LENGTH; i++) {
		pthread_mutex_lock( &mutex);
		queue[queue_len] = i;
		queue_len++;
		pthread_cond_signal( &cv);
		pthread_mutex_unlock( &mutex);
		sleep( 1);
	}
}

void *consumer( void *param)
{
	int seconds = 0;
	int e;
	int i;
	int item;
	struct timespec now;

	for (i = 0; i < QUEUE_LENGTH; i++) {

		seconds = 0;	// reset

		e = pthread_mutex_lock( &mutex);
		if (e) {
			fprintf( stderr, "pthread_mutex_lock() failed: E=%d[ %s ]\n", e, strerror( e));
			continue;
		}

		while (queue_len == 0) {
			now.tv_sec = time( NULL) + 1;
			now.tv_nsec = 0;
			e = pthread_cond_timedwait( &cv, &mutex, &now);
			if (e) {
				seconds++;
			}
		}

		item = queue[queue_len];
		queue_len--;
		if (seconds) {
			printf( "%d seconds waited\n", seconds);
		}

		e = pthread_mutex_unlock( &mutex);
		if (e) {
			fprintf( stderr, "pthread_mutex_unlock() failed: E=%d[ %s ]\n", e, strerror( e));
			continue;
		}

		printf( "Received %d\n", item);
	}

	return NULL;
}

int main( int argc, char **argv)
{
	pthread_t threads[2];

	pthread_cond_init( &cv, 0);
	pthread_mutex_init( &mutex, 0);
	pthread_create( &threads[0], 0, producer, 0);
	pthread_create( &threads[1], 0, consumer, 0);
	pthread_join( threads[0], 0);
	pthread_join( threads[1], 0);
	//pthread_mutex_destory( &mutex);
	//pthread_cond_destory( &cv);

	return 0;
}

