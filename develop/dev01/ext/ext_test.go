package ext

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	GetTime("")
	fmt.Printf("\n")

	time.Sleep(time.Second * 5)

	GetTime("0.beevik-ntp.pool.ntp.org")
	fmt.Printf("\n")
}
