package example

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/RobertGumpert/gopherpc"
)

type Data struct {
	Params string `json:"params"`
}

func TestSimpleFlow(t *testing.T) {
	//
	// Client -->
	//
	dataForRequest := &Data{
		Params: "hello, world!",
	}
	reqBody := &gopherpc.Request{
		Jsonrpc: gopherpc.ProtoVersion,
		Method: "hello",
		Params: dataForRequest,
		ID: 1,
	}
	bts, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}
	//
	// Server -->
	//
	request, response, err := gopherpc.ParseRequest(bts)
	if err != nil {
		t.Fatal(err)
	}
	response, err = request.ParseParams(new(Data))
	if err != nil {
		t.Fatal(err)
	}
	request.Params.(*Data).Params = "Hello, from server!"
	response, err = request.Response(request.Params)
	if err != nil {
		t.Fatal(err)
	}
	bts, err := response.Marshall()
	if err != nil {
		t.Fatal(err)
	}
	//
	// Client -->
	//
	is := gopherpc.IsResponse(bts)
	if is {
		response, err := gopherpc.ParseResponse(bts)
		if err != nil {
			t.Fatal(err)
		}
		err = response.ParseResult(new(Data))
		log.Println(response.Result.(*Data).Params)
	} else {
		t.Fatal("Non response")
	}
}