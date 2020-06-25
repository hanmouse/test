package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/heptiolabs/healthcheck"
)

var testURL = "https://www.google.com"

func main() {

	var isServerDone chan bool

	asyncCheckInterval := 3 * time.Second
	//timeout := 5 * time.Microsecond

	health := healthcheck.NewHandler()

	//health.AddLivenessCheck("liveness-check", healthcheck.DNSResolveCheck(testURL, timeout))
	//health.AddLivenessCheck("liveness-check", healthcheck.HTTPGetCheck(testURL, timeout))
	//health.AddLivenessCheck("liveness-check", healthcheck.Async(healthcheck.HTTPGetCheck(testURL, timeout), asyncCheckInterval))
	//health.AddLivenessCheck("liveness-check", fakeCheckLiveness)

	// Our app is not ready if we can't connect to our database (`var db *sql.DB`) in <1s.
	//db := sql.DB
	//health.AddReadinessCheck("database", healthcheck.DatabasePingCheck(db, 1*time.Second))

	health.AddLivenessCheck("liveness-check", healthcheck.Async(checkLiveness, asyncCheckInterval))
	health.AddReadinessCheck("readiness-check", healthcheck.Async(checkReadiness, asyncCheckInterval))

	time.Sleep(1 * time.Second)

	go http.ListenAndServe("0.0.0.0:8080", health)

	<-isServerDone
}

func checkLiveness() error {
	fmt.Println("I'm alive!!")
	return nil
}

func checkReadiness() error {

	var err error

	fileToCheck := "./testfile"

	_, err = os.Stat(fileToCheck)
	if os.IsExist(err) {
		err = nil
	}

	return err
}
