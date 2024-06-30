package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Println(os.Stderr, "Error connecting", err)
		os.Exit(1)
	}

	currentTime := time.Now()

	fmt.Println("NTP time: ", ntpTime)
	fmt.Println("Local time: ", currentTime)

}
