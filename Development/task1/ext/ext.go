package ext

// package exactTime

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

// GetTime : show current Time or return error and exit status
func GetTime(host string) {
	if host == "" {
		host = "0.beevik-ntp.pool.ntp.org"
	}

	response, err := ntp.Query(host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error message: %s\n", err)
		os.Exit(1)
	}

	currentTime := time.Now()
	exactTime := time.Now().Add(response.ClockOffset)

	fmt.Printf("%s\n", currentTime)
	fmt.Printf("%s\n", exactTime)
}
