#include "reader.h"
#include "file.h"

static reader_t readers[NUM_READERS];

void *read_msgs( void *param)
{
	int thread_id = (int) param;
	int e;
	unsigned int count = 0U;
	unsigned int int_data;
	unsigned char data[MAX_QUEUE_DATA_SIZE];
	size_t data_size;
	queue_t *queue = &readers[thread_id].queue;
	file_info_t *file = &readers[thread_id].file;

	printf( "thread %d created\n", thread_id);

	e = init_queue( queue, MAX_QUEUE_SIZE);
	if (e < 0) {
		fprintf( stderr, "init_queue() failed\n");
		return NULL;
	}

	e = init_file( file, thread_id);
	if (e < 0) {
		fprintf( stderr, "init_file() failed\n");
		return NULL;
	}

	while (should_thread_keep_alive() && (count < MAX_QUEUE_SIZE)) {
		e = get_queue_data( queue, data, &data_size);
		if (e < 0) {
			fprintf( stderr, "get_queue_data() failed\n");
			continue;
		} else {
			if (data_size == 0) {
				continue;
			}
		}

		memcpy( &int_data, data, data_size);
		printf( "[thread %d] data=%u, data_size=%zu\n", thread_id, int_data, data_size);

		fprintf( file->fp, "data=%d\n", int_data);
		count++;
		printf( "[thread %d] count=%u\n", thread_id, count);
	}

	finish_file( file->fp);

	return NULL;
}

reader_t *get_readers( void)
{
	return readers;
}

