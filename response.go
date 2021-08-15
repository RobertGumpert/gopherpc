package gopherpc

import (
	"encoding/json"
	"regexp"
)

type IResponse interface {
	Marshall() ([]byte, error)
	String() (string, error)
}

type response struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	ID      interface{} `json:"id"`
	//
	bytesRepresentation  []byte `json:"-"`
	stringRepresentation string `json:"-"`
}

func (this *response) Marshall() ([]byte, error) {
	if this.bytesRepresentation == nil ||
		len(this.bytesRepresentation) == 0 {
		bts, err := json.Marshal(this)
		if err != nil {
			return nil, err
		}
		this.bytesRepresentation = bts
	}
	return this.bytesRepresentation, nil
}

func (this *response) String() (string, error) {
	if this.stringRepresentation == "" {
		_, err := this.Marshall()
		if err != nil {
			return "", err
		}
		this.stringRepresentation = string(this.bytesRepresentation)
	}
	return this.stringRepresentation, nil
}

func IsResponse(bts []byte) bool {
	isError, _ := regexp.Match(
		errorRegexString,
		bts,
	)
	return !isError
}

func ParseResponse(bts []byte) (*response, error) {
	var (
		resp = new(response)
	)
	err := json.Unmarshal(bts, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (this *response) ParseResult(userTypeResult interface{}) error {
	if resultBytes, err := json.Marshal(this.Result); err != nil {
		return err
	} else {
		if this.Result != nil {
			if err := json.Unmarshal([]byte(resultBytes), userTypeResult); err != nil {
				return err
			} else {
				this.Result = userTypeResult
			}
		}
	}
	return nil
}
