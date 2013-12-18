#include "thread_life.h"

static int keepalive = true;

int should_thread_keep_alive( void)
{
	return keepalive;
}

void terminate_thread( void)
{
	keepalive = false;
}

