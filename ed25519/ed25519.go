//
// ed25519.go
//
package ed25519

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"

	mySignature "github.com/k4k3ru-hub/signature-go"
)


//
// Generate a key pair as Base64.
//
func GenerateKeyPairBase64() (publicKey string, privateKey string, err error) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return
	}
	publicKey = base64.StdEncoding.EncodeToString(pub)
	privateKey = base64.StdEncoding.EncodeToString(priv)
	return
}


//
// Sign a canonicalized JSON data using Ed25519 and returns a Base64 encoded signature.
//
func SignJson[T any](data T, privateKeyBase64 string) (string, error) {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(privateKeyBase64)
	if err != nil {
		return "", err
	}
	if len(privateKeyBytes) != ed25519.PrivateKeySize {
		return "", fmt.Errorf("Invalid private key length.\n")
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	canonMsg, err := mySignature.CanonicalizeJSON(jsonBytes)
	if err != nil {
		return "", err
	}

	sig := ed25519.Sign(privateKeyBytes, canonMsg)
	return base64.StdEncoding.EncodeToString(sig), nil
}


//
// Verify a canonicalized JSON data using Ed25519.
//
func VerifyJson[T any](data T, publicKeyBase64 string, signatureBase64 string) (bool, error) {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyBase64)
	if err != nil {
		return false, fmt.Errorf("failed to decode public key: %w", err)
	}
	if len(publicKeyBytes) != ed25519.PublicKeySize {
		return false, fmt.Errorf("invalid public key length")
	}

	sigBytes, err := base64.StdEncoding.DecodeString(signatureBase64)
	if err != nil {
		return false, fmt.Errorf("failed to decode signature: %w", err)
	}

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return false, fmt.Errorf("failed to marshal struct: %w", err)
	}

	canonMsg, err := mySignature.CanonicalizeJSON(jsonBytes)
	if err != nil {
		return false, fmt.Errorf("failed to canonicalize JSON: %w", err)
	}

	valid := ed25519.Verify(publicKeyBytes, canonMsg, sigBytes)
	return valid, nil
}

