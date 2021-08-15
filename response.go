package gopherpc

import (
	"encoding/json"
	"regexp"
)


type IResponse interface {
	Marshall() ([]byte, error)
}

type response struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	ID      interface{} `json:"id"`
}

func (this *response) Marshall() ([]byte, error) {
	return json.Marshal(this)
}

func IsResponse(bts []byte) bool {
	isError, _ := regexp.Match(
		errorRegexString,
		bts,
	)
	return !isError
}

func ParseResponse(bts []byte, userTypeResult ...interface{}) (*response, error) {
	var (
		resp = new(response)
	)
	err := json.Unmarshal(bts, resp)
	if err != nil {
		return nil, err
	}
	if resultBytes, err := json.Marshal(resp.Result); err != nil {
		return nil, err
	} else {
		if resp.Result != nil {
			if err := json.Unmarshal([]byte(resultBytes), userTypeResult[0]); err != nil {
				return nil, err
			} else {
				resp.Result = userTypeResult[0]
			}
		}
	}
	return resp, nil
}