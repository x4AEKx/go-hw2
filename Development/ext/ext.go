package ext

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func GetTime(host string) {
	if host == "" {
		host = "0.beevik-ntp.pool.ntp.org"
	}

	response, err := ntp.Query(host)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error message: %s\n", err)
		os.Exit(1)
	}

	time := time.Now().Add(response.ClockOffset)

	fmt.Printf("%s", time)
}
