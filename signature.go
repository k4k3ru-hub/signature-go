//
// signature.go
//
package signature

import (
	"bytes"
	"encoding/json"
	"sort"
	"time"

	
)


//
// Canonicalize JSON bytes.
// Returns a compact, key-sorted version.
// This is useful for generating consistent signatures over structured data.
//
func CanonicalizeJSON(data interface{}) ([]byte, error) {
	var normalized interface{}
	switch v := data.(type) {
	case string:
		if err := json.Unmarshal([]byte(v), &normalized); err != nil {
			return nil, err
		}
	case []byte:
		if err := json.Unmarshal(v, &normalized); err != nil {
			return nil, err
		}
	default:
		jsonBytes, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(jsonBytes, &normalized); err != nil {
			return nil, err
		}
	}

	var buf bytes.Buffer
	if err := encodeCanonical(&buf, normalized); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}


//
// Verify UNIX timestamp.
//
func VerifyTimestamp(unixTimestamp int64, tolerance time.Duration) bool {
	now := time.Now().Unix()
	upper := now + int64(tolerance.Seconds())
	lower := now - int64(tolerance.Seconds())
	return lower <= unixTimestamp && unixTimestamp <= upper
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
