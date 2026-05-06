package utils

import (
	"encoding/json"
	"io"
)

func ResponseDecoderHandler(content io.Reader, v any) error {
	decoder := json.NewDecoder(content)
	decoder.DisallowUnknownFields()
	return decoder.Decode(v)
}
