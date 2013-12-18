#include "queue.h"

static int queue_is_empty( const queue_t *queue);
static int queue_is_full( const queue_t *queue);

int init_queue( queue_t *queue, size_t size)
{
	queue->size = size + 1;	/* include empty element */
	queue->front = queue->rear = 0;
	queue->data = calloc( queue->size, MAX_QUEUE_DATA_SIZE);
	if (queue->data == NULL) {
		fprintf( stderr, "calloc() failed: size=%zu, E=%d[ %s ]\n", queue->size * MAX_QUEUE_DATA_SIZE, errno, strerror( errno));
		return -1;
	}

	return 0;
}

static int queue_is_full( const queue_t *queue)
{
	return (((queue->rear + 1) % queue->size) == queue->front);
}

int insert_queue_data( queue_t *queue, const unsigned char *data, size_t data_size)
{
	if (queue_is_full( queue)) {
		fprintf( stderr, "queue is full\n");
		return -1;
	}

	if (data_size > MAX_QUEUE_DATA_SIZE) {
		fprintf( stderr, "data is too big: data_size=%zu, max_size=%zu\n", data_size, MAX_QUEUE_DATA_SIZE);
		return -1;
	}

	memcpy( &queue->data[queue->rear * MAX_QUEUE_DATA_SIZE], data, data_size);
	queue->data_size[queue->rear] = data_size;
	queue->rear = (queue->rear + 1) % queue->size;

	return 0;
}

static int queue_is_empty( const queue_t *queue)
{
	return (queue->rear == queue->front);
}

int get_queue_data( queue_t *queue, unsigned char *data, size_t *data_size)
{
	*data_size = 0;

	if (queue_is_empty( queue)) {
		return 0;
	}

	memcpy( data, &queue->data[queue->front * MAX_QUEUE_DATA_SIZE], queue->data_size[queue->front]);
	*data_size = queue->data_size[queue->front];
	queue->front = (queue->front + 1) % queue->size;

	return 0;
}

