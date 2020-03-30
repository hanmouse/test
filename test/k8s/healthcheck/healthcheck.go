package main

import (
	"time"

	"github.com/heptiolabs/healthcheck"
)

// TODO
func main() {
	health := healthcheck.NewHandler()

	// Our app is not happy if we've got more than 100 goroutines running.
	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))

	// Our app is not ready if we can't resolve our upstream dependency in DNS.
	health.AddReadinessCheck(
		"upstream-dep-dns",
		healthcheck.DNSResolveCheck("upstream.example.com", 50*time.Millisecond))

	// Our app is not ready if we can't connect to our database (`var db *sql.DB`) in <1s.
	//db := sql.DB
	//health.AddReadinessCheck("database", healthcheck.DatabasePingCheck(db, 1*time.Second))
}
