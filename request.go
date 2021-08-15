package gopherpc

import (
	"encoding/json"
)

const (
	ProtoVersion string = "2.0"
)

type Request struct {
	Jsonrpc        string      `json:"jsonrpc"`
	Method         string      `json:"method"`
	Params         interface{} `json:"params"`
	ID             interface{} `json:"id"`
	isNotification bool
}

func ParseRequest(bts []byte, userTypeParams ...interface{}) (*Request, IResponse, error) {
	var (
		request = new(Request)
	)
	if err := json.Unmarshal(bts, request); err != nil {
		return nil, Error(ErrParse, err.Error()), err
	}
	if paramsBytes, err := json.Marshal(request.Params); err != nil {
		return nil, Error(ErrParse, err.Error()), err
	} else {
		if request.Params != nil {
			if err := json.Unmarshal([]byte(paramsBytes), userTypeParams[0]); err != nil {
				return nil, Error(ErrParse, err.Error()), err
			} else {
				request.Params = userTypeParams[0]
			}
		}
	}
	if request.ID == nil {
		request.isNotification = true
	} else {
		request.isNotification = false
	}
	return request, nil, nil
}

func (this *Request) Response(result interface{}) (IResponse, error) {
	if this.isNotification {
		return nil, NotificationHasntResponse
	}
	return &response{
		Jsonrpc: ProtoVersion,
		Result:  result,
		ID:      this.ID,
	}, nil
}

func (this *Request) Error(code ErrCode, message string) (IResponse, error) {
	if this.isNotification {
		return nil, NotificationHasntResponse
	}
	return Error(code, message), nil
}
