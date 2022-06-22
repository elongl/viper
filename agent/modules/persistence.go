//go:build !windows

package modules

import (
	"log"
)

func EnsurePersistence() error {
	log.Printf("Persistence is currently unavailable on non-Windows agents.")
	return nil
}
