package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func StartSolBot() (string, error) {
	fmt.Println("Starting Solana Bot...")
	return "👋 *Welcome to Solana Bot!\nCreated by Jojosaturo.\n\nClick the button below to generate your Solana wallet address.", nil
}

func CreateSolAccount() (string, error) {
	fmt.Println("Creating Solana Account...")

	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	address := "sol_" + hex.EncodeToString(randomBytes[:])

	return fmt.Sprintf("✅ *Solana Wallet Created!*\n\n`%s`", address), nil
}
