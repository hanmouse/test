package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/heptiolabs/healthcheck"
)

var testURL = "https://heptio.com"

// TODO
func dumpRequest(handler healthcheck.Handler, httpMethod string, location string) string {
	resp, err := http.Get(testURL)
	if err != nil {
		return fmt.Sprint(err)
	}

	defer resp.Body.Close()

	//return fmt.Sprintf("%#v\n", resp.Header)
	return fmt.Sprintf("Status: %#v\n", resp.Status)

	/*
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Sprint(err)
		}

		return string(body)
	*/
}

// TODO
func main() {

	//var isServerDone chan bool

	health := healthcheck.NewHandler()

	health.AddLivenessCheck("liveness-check", healthcheck.HTTPGetCheck(testURL, 5*time.Second))

	go http.ListenAndServe("0.0.0.0:8080", health)

	//<-isServerDone

	fmt.Print(dumpRequest(health, "GET", "/live"))

	/*
		health.AddReadinessCheck(
			"upstream-dep-dns",
			healthcheck.DNSResolveCheck("upstream.example.com", 50*time.Millisecond))
	*/

	// Our app is not ready if we can't connect to our database (`var db *sql.DB`) in <1s.
	//db := sql.DB
	//health.AddReadinessCheck("database", healthcheck.DatabasePingCheck(db, 1*time.Second))
}
