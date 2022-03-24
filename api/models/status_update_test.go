package models

import (
	"testing"
)

func Test_ValidateStatusUpdate(t *testing.T) {
	su := &StatusUpdate{
		Asset_Status: "Escrow",
		Mint_Addr:    "0x_fakemintaddr",
	}

	err := su.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

func Test_MissingMintAddrFails(t *testing.T) {
	su := &StatusUpdate{
		Asset_Status: "Escrow",
	}

	err := su.Validate()

	if err == nil {
		t.Fatal(err)
	}
}
