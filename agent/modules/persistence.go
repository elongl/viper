//go:build !windows

package modules

import (
	"log"
)

func EnsurePersistence() error {
	log.Print("Persistence is currently unavailable on non-Windows agents.")
	return nil
}
