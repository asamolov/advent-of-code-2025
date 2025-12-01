package utils

import (
	"fmt"
	"time"
)

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("[%s]: %s\n", name, time.Since(start))
	}
}
