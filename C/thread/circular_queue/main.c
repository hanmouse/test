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
#include "cirqueue_test.h"

/* main thread is the "message sender" */
int main( int argc, char **argv)
{
	int i, j;
	int e;
	unsigned int tmpval = 0;
	unsigned char data[MAX_QUEUE_DATA_SIZE];
	reader_t *readers = get_readers();

	for (i = 0; i < NUM_READERS; i++) {
		e = pthread_create( &readers[i].thread, NULL, read_msgs, (void *) i);
		if (e) {
			fprintf( stderr, "pthread_create() failed: thread_id=%d, E=%d[ %s ]\n", i, e, strerror( e));
			return 1;
		}
	}

	sleep( 2);

	for (i = 0; i < NUM_READERS; i++) {
		for (j = 0; j < MAX_QUEUE_SIZE; j++) {
			memcpy( data, &tmpval, sizeof( tmpval));
			e = insert_queue_data( &readers[i].queue, data, sizeof( tmpval));
			if (e < 0) {
				fprintf( stderr, "insert_queue_data() failed: thread_id=%d\n", i);
				continue;
			}
			printf( "data inserted into thread %d: data=%u\n", i, tmpval);
			tmpval++;
		}
	}

	for (i = 0; i < NUM_READERS; i++) {
		e = pthread_join( readers[i].thread, NULL);
		if (e) {
			fprintf( stderr, "pthread_join() failed: thread_id=%d, E=%d[ %s ]\n", i, e, strerror( e));
			continue;
		}
	}

	return 0;
}

