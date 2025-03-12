package wallet

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

func GenerateEthereumWallet() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	publicKey := privateKey.Public().(*ecdsa.PublicKey)

	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	address := crypto.Keccak256(publicKeyBytes[1:])[12:]

	return crypto.PubkeyToAddress(*publicKey).Hex(), hex.EncodeToString(address), nil
}
