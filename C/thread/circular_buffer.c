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

#define BUF_SIZE	16

volatile int add_here, remove_here;
volatile int volatile buffer[BUF_SIZE];

void clear_buffer( void)
{
	memset( buffer, 0, sizeof( buffer));
}

int next( int current)
{
	return (current + 1) & (BUF_SIZE - 1);
}

void add_to_buffer( int value)
{
	while (buffer[next( add_here)]) {}
	buffer[add_here] = value;
	add_here = next( add_here);
}

int remove_from_buffer( void)
{
	int value;

	while (value = buffer[remove_here] == 0) {}
	buffer[remove_here] = 0;
	remove_here = next( remove_here);
	return value;
}

void *producer( void *param)
{
	int i;

	for (i = 1; i < 10000000; i++) {
		add_to_buffer( i);
	}
}

void *consumer( void *param)
{
	while (remove_from_buffer() != 9999999) {}
}

int main( int argc, char **argv)
{
	pthread_t threads[2];

	clear_buffer();

	pthread_create( &threads[0], NULL, producer, 0);
	pthread_create( &threads[1], NULL, consumer, 0);
	pthread_join( &threads[0], NULL);
	pthread_join( &threads[1], NULL);

	return 0;
}

