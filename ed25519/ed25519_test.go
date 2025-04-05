//
// ed25519_test.go
//
package ed25519

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"testing"
)


//
// Test for generate key pair.
//
func TestGenerateKeyPairBase64(t *testing.T) {
	pub, priv, err := GenerateKeyPairBase64()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	pubBytes, err := base64.StdEncoding.DecodeString(pub)
	if err != nil {
		t.Fatalf("Invalid base64 public key: %v", err)
	}
	if len(pubBytes) != ed25519.PublicKeySize {
		t.Errorf("Unexpected public key size: %d", len(pubBytes))
	}

	privBytes, err := base64.StdEncoding.DecodeString(priv)
	if err != nil {
		t.Fatalf("Invalid base64 private key: %v", err)
	}
	if len(privBytes) != ed25519.PrivateKeySize {
		t.Errorf("Unexpected private key size: %d", len(privBytes))
	}
}


//
// Test for signing and verifying JSON.
//
func TestSignAndVerifyJson(t *testing.T) {
	pubStr, privStr, err := GenerateKeyPairBase64()
	if err != nil {
		t.Fatalf("Failed to generate keys: %v", err)
	}

	data := struct {
        AccountID string `json:"accountId"`
        Timestamp int64  `json:"timestamp"`
        Params struct{
            Symbol string `json:"symbol"`
        } `json:"params"`
        Arrays []struct{
            Side string `json:"side"`
            Amount string `json:"amount"`
        } `json:"orders"`
    }{
        AccountID: "user123",
        Timestamp: 1712000000000,
        Params: struct{
            Symbol string `json:"symbol"`
        }{
            Symbol: "BTC/USDT",
        },
        Arrays: []struct{
            Side string `json:"side"`
            Amount string `json:"amount"`
        }{
            {Side: "BUY", Amount: "0.1"},
            {Side: "SELL", Amount: "0.2"},
        },
    }
	jsonBytes, _ := json.Marshal(data)

	signature, err := SignJson(jsonBytes, privStr)
	if err != nil {
		t.Fatalf("Failed to sign: %v", err)
	}
	t.Logf("Signed JSON. (signature: %s)\n", string(signature))

	ok, err := VerifyJson(jsonBytes, pubStr, signature)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if !ok {
		t.Error("Failed to verify data.")
	}	
}

