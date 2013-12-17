#include "file.h"

int init_file( file_info_t *file, int thread_id)
{
	sprintf( file->path, "./output_thread%d.txt", thread_id);

	file->fp = fopen( file->path, "w");
	if (file->fp == NULL) {
		fprintf( stderr, "fopen() failed: file=\"%s\", E=%d[ %s ]\n", file->path, errno, strerror( errno));
		return -1;
	}

	return 0;
}

