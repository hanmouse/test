#ifndef FILE_H
#define FILE_H

#include <stdio.h>
#include <string.h>
#include <errno.h>

typedef struct file_info_s {
	char path[128];
	FILE *fp;
} file_info_t;

int init_file( file_info_t *file, int thread_id);

#endif /* FILE_H */

