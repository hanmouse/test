#ifndef QUEUE_H
#define QUEUE_H

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <errno.h>
#include <stdbool.h>

#define MAX_QUEUE_SIZE		128
#define MAX_QUEUE_DATA_SIZE	64

/* circular queue */
typedef struct queue_s {
	size_t size;
	int front;
	int rear;
	unsigned char *data;	/* vector of elements */
	size_t data_size[MAX_QUEUE_SIZE];
} queue_t;

int get_queue_data( queue_t *queue, unsigned char *data, size_t *data_size);
int init_queue( queue_t *queue, size_t size);
int insert_queue_data( queue_t *queue, const unsigned char *data, size_t data_size);

#endif /* QUEUE_H */

