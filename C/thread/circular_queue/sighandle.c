#include "sighandle.h"
#include "thread_life.h"

static void sighandler_to_ignore( int signo)
{
	fprintf( stderr, "caught signal %d\n", signo);
}

static void sighandler_to_term( int signo)
{
	fprintf( stderr, "caught signal %d\n", signo);
	terminate_thread();
}

static int set_sigaction( struct sigaction *sigact, void (*handler)(int), int n_signals, ...)
{
	int i;
	int signo;
	va_list ap;

	sigact->sa_handler = handler;
	sigemptyset( &sigact->sa_mask);
	sigact->sa_flags = 0;

	va_start( ap, n_signals);

	for (i = 0; i < n_signals; i++) {
		signo = va_arg( ap, int);
		if (sigaction( signo, sigact, 0) < 0) {
			fprintf( stderr, "sigaction() failed: E=%d[ %s ], signo=%d\n", errno, strerror( errno), signo);
			return -1;
		}
	}

	va_end( ap);

	return 0;
}

int init_sig_handler( void)
{
	int e;
	struct sigaction sigact_to_ignore;
	struct sigaction sigact_to_term;

	e = set_sigaction( &sigact_to_ignore, sighandler_to_ignore, 2, SIGPIPE, SIGALRM);
	if (e < 0) {
		fprintf( stderr, "set_sigaction() failed\n");
		return -1;
	}

	e = set_sigaction( &sigact_to_term, sighandler_to_term, 3, SIGTERM, SIGQUIT, SIGINT);
	if (e < 0) {
		fprintf( stderr, "set_sigaction() failed\n");
		return -1;
	}

	return 0;
}

