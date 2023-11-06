//go:build mage
// +build mage

package main

import (
	"fmt"

	"sigs.k8s.io/release-utils/mage"
)

var Default = Verify

func Verify() error {
	fmt.Println("Running external dependency checks...")
	if err := mage.VerifyDeps("v0.4.1", "", "", true); err != nil {
		return err
	}
	return nil
}
