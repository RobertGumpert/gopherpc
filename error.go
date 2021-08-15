package gopherpc

import "encoding/json"

type errorContext struct {
	Code    ErrCode `json:"code"`
	Message string  `json:"message"`
}

type errorResponse struct {
	Jsonrpc string        `json:"jsonrpc"`
	Error   *errorContext `json:"error"`
	ID      interface{}   `json:"id"`
}

func (this *errorResponse) Marshall() ([]byte, error) {
	return json.Marshal(this)
}

func Error(code ErrCode, message string) IResponse {
	this := &errorResponse{
		Jsonrpc: "2.0",
		Error: &errorContext{
			Code:    code,
			Message: message,
		},
		ID: nil,
	}
	return this
}

func ParseError(bts []byte) (*errorResponse, error) {
	var (
		resp = new(errorResponse)
	)
	err := json.Unmarshal(bts, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}