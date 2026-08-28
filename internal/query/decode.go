package query

import "encoding/json"

func storageDecode(d []byte, v any) error { return json.Unmarshal(d, v) }
