package ext

import (
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	GetTime("")

	time.Sleep(time.Second * 5)

	GetTime("0.beevik-ntp.pool.ntp.org")
}
