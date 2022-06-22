//go:build !windows

package modules

import (
	"log"
)

func EnsurePersistence() error {
	log.Printf("persistence is currently unavailable on non-Windows agents")
	return nil
}
