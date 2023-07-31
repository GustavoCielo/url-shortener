package json

import (
	"encoding/json"

	"github.com/GustavoCielo/url-shortener/shortener"
	"github.com/pkg/errors"
)

type Redirect struct{}

// Decode is used to decode the json data
func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := json.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

// Encode is used to encode the json data
func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Enconde")
	}
	return rawMsg, nil
}
