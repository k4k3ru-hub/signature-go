//
// signature.go
//
package signature

import (
	"bytes"
	"encoding/json"
	"sort"
)


//
// Canonicalize JSON bytes.
// Returns a compact, key-sorted version.
// This is useful for generating consistent signatures over structured data.
//
func CanonicalizeJSON(jsonBytes []byte) ([]byte, error) {
	var data interface{}
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if err := encodeCanonical(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}


//
// Encodes for canonicalized JSON data.
//
func encodeCanonical(buf *bytes.Buffer, v interface{}) error {
	switch val := v.(type) {
	case map[string]interface{}:
		buf.WriteByte('{')
		keys := make([]string, 0, len(val))
		for k := range val {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for i, k := range keys {
			if i > 0 {
				buf.WriteByte(',')
			}
			keyBytes, _ := json.Marshal(k)
			buf.Write(keyBytes)
			buf.WriteByte(':')
			if err := encodeCanonical(buf, val[k]); err != nil {
				return err
			}
		}
		buf.WriteByte('}')
	case []interface{}:
		buf.WriteByte('[')
		for i, item := range val {
			if i > 0 {
				buf.WriteByte(',')
			}
			if err := encodeCanonical(buf, item); err != nil {
				return err
			}
		}
		buf.WriteByte(']')
	default:
		enc, err := json.Marshal(val)
		if err != nil {
			return err
		}
		buf.Write(enc)
	}
	return nil
}
