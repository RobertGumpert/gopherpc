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
	//
	bytesRepresentation  []byte `json:"-"`
	stringRepresentation string `json:"-"`
}

func (this *errorResponse) Marshall() ([]byte, error) {
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

func (this *errorResponse) String() (string, error) {
	if this.stringRepresentation == "" {
		_, err := this.Marshall()
		if err != nil {
			return "", err
		}
		this.stringRepresentation = string(this.bytesRepresentation)
	}
	return this.stringRepresentation, nil
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
