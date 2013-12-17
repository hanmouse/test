#ifndef READER_H
#define READER_H

#include <stdio.h>
#include <pthread.h>
#include "queue.h"
#include "file.h"

#define NUM_READERS	4

typedef struct reader_s {
	pthread_t thread;
	queue_t queue;
	file_info_t file;
} reader_t;

reader_t *get_readers( void);
void *read_msgs( void *param);

#endif /* READER_H */

