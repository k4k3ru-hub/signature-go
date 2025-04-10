//
// signature_test.go
//
package signature

import (
	"encoding/json"
	"testing"
)


//
// Test for JSON string.
//
func TestCanonicalizeJsonForJsonString(t *testing.T) {
	jsonStr := `{"b":2,"a":1}`
	expected := `{"a":1,"b":2}`

	out, err := CanonicalizeJSON([]byte(jsonStr))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if string(out) != expected {
		t.Errorf("Unexpected value. (expected: %s, got: %s)", expected, string(out))
	}
}


//
// Test for JSON struct.
//
func TestCanonicalizeJsonForJsonStruct(t *testing.T) {
	data := struct {
		AccountID string `json:"accountId"`
		Timestamp int64  `json:"timestamp"`
	}{
		AccountID: "user123",
		Timestamp: 1712000000000,
	}
	jsonBytes, _ := json.Marshal(data)
	expected := `{"accountId":"user123","timestamp":1712000000000}`

	out, err := CanonicalizeJSON(jsonBytes)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }

    if string(out) != expected {
        t.Errorf("Unexpected value. (expected: %s, got: %s)", expected, string(out))
    }
}


//
// Test for JSON struct nested parameters.
//
func TestCanonicalizeJsonForJsonStructNestedParams(t *testing.T) {
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
	expected := `{"accountId":"user123","orders":[{"amount":"0.1","side":"BUY"},{"amount":"0.2","side":"SELL"}],"params":{"symbol":"BTC/USDT"},"timestamp":1712000000000}`

	out, err := CanonicalizeJSON(jsonBytes)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }

    if string(out) != expected {
		t.Errorf("Unexpected value. (expected: %s, got: %s)", expected, string(out))
    }
}


//
// Test for map object nested parameters.
//
func TestCanonicalizeJsonForMapNestedParams(t *testing.T) {
	data := map[string]interface{}{
		"accountId": "user123",
		"timestamp": 1712000000000,
		"params": map[string]interface{}{
			"symbol": "BTC/USDT",
		},
		"orders": []map[string]interface{}{
			{"side": "BUY", "amount": "0.1"},
			{"side": "SELL", "amount": "0.2"},
			{"side": "BUY", "amount": "0.3"},
		},
	}
	jsonBytes, _ := json.Marshal(data)
	expected := `{"accountId":"user123","orders":[{"amount":"0.1","side":"BUY"},{"amount":"0.2","side":"SELL"},{"amount":"0.3","side":"BUY"}],"params":{"symbol":"BTC/USDT"},"timestamp":1712000000000}`

	out, err := CanonicalizeJSON(jsonBytes)
    if err != nil {
        t.Fatalf("Unexpected error: %v", err)
    }
    if string(out) != expected {
		t.Errorf("Unexpected value. (expected: %s, got: %s)", expected, string(out))
    }
	t.Logf("Canonicalized JSON: %s.\n", string(jsonBytes))
}
