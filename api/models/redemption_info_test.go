package models

import (
	"log"
	"testing"
)

func Test_ValidateRedemptionInfo(t *testing.T) {
	ri := &RedemptionInfo{
		Wallet_PK:                    "0x_fakewalletpk",
		Redemption_Info_Account_Addr: "0x_fakeredemptioninfoaccountpk",
		Baxus_Escrow_Addr:            "0x_fakebaxusescrowpk",
		Mint_Addr:                    "0x_fakemintaddr",
	}

	err := ri.Validate()
	if err != nil {
		log.Fatal(t)
	}
}

func Test_MissingRedemptionInfoPK(t *testing.T) {
	ri := &RedemptionInfo{
		Wallet_PK:         "0x_fakewalletpk",
		Baxus_Escrow_Addr: "0x_fakebaxusescrowpk",
		Mint_Addr:         "0x_fakemintaddr",
	}

	err := ri.Validate()
	if err == nil {
		log.Fatal(t)
	}
}

func Test_MissingBaxusEscrowPK(t *testing.T) {
	ri := &RedemptionInfo{
		Wallet_PK:                    "0x_fakewalletpk",
		Redemption_Info_Account_Addr: "0x_fakeredemptioninfoaccountpk",
		Mint_Addr:                    "0x_fakemintaddr",
	}

	err := ri.Validate()
	if err == nil {
		log.Fatal(t)
	}
}

func Test_MissingMintAddr(t *testing.T) {
	ri := &RedemptionInfo{
		Wallet_PK:                    "0x_fakewalletpk",
		Redemption_Info_Account_Addr: "0x_fakeredemptioninfoaccountpk",
		Baxus_Escrow_Addr:            "0x_fakebaxusescrowpk",
	}

	err := ri.Validate()
	if err == nil {
		log.Fatal(t)
	}
}
