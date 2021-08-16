# Parse Request.
## Struct

```go
type Request struct {
	Jsonrpc        string      `json:"jsonrpc"`
	Method         string      `json:"method"`
	Params         interface{} `json:"params"`
	ID             interface{} `json:"id"`
}
```

## Description

After you have received the byte array containing the body, pass them as an argument. If the parsing was successful, then the function will return you only the "request" variable, otherwise it will return the "response" variable already in the error representation format according to the JSON-RPC standard, and also return an error. After that, call the parsing method of the value that must be represented by the "parameters" key, passing a pointer to the object as an argument.

```go

//
// Server
//

type Data struct {
    ...
}


...


request, response, err := gopherpc.ParseRequest(bts)    
if err != nil {
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
response, err = request.ParseParams(new(Data))
if err != nil {
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}


...


switch request.Method {
    case "method":

        ...
        
    default:
        response, err := request.Error(gopherpc.ErrMethodNotFound, "Oups")
        if err != nil {
            ...
        }

        ...

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
}

```

In order to return a response, use the "Response" method of the "Request" object, passing it as an argument a pointer to an object that represents the value by the "result" key. If you want to return an error, use the "Error" method of the "Request" object, passing it an error code as the first argument, and a string describing the error as the second argument. Note that, according to the standard, if the request was a notification, then both methods will return an error.

```go

//
// Server
//

type Some struct {
    Hello string `json:"hello"`
}


...


response, err := request.Response(&Some{Hello: "World"})
if err != nil {
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

...

response, err := request.Error(gopherpc.ErrMethodNotFound, "Hello world")
if err != nil {
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

...

bts, err = response.Marshall()
if err != nil {
	response = gopherpc.Error(gopherpc.ErrInternalError, err.Error())
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
} else {
	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
    go func() {
        err := rabbitMQChannel.Publish(
            "",    
            q.Name,
            false, 
            false,
            amqp.Publishing{
                ContentType: "text/plain",
                Body:        bts,
            },
        )
        ...
    }()
    
}

```

# Parse Response.
## Struct

```go

type IResponse interface {
	Marshall() ([]byte, error)
	String() (string, error)
}

type response struct {
	Jsonrpc string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	ID      interface{} `json:"id"`
}

type errorContext struct {
	Code    ErrCode `json:"code"`
	Message string  `json:"message"`
}

type errorResponse struct {
	Jsonrpc string        `json:"jsonrpc"`
	Error   *errorContext `json:"error"`
	ID      interface{}   `json:"id"`
}
```

## Description
Before parsing the answer, you can use the "IsResponse()" function to find out if the answer is a normal answer or an error.

```go

//
// Client
//

type Data struct {
    ...
}


...


is := gopherpc.IsResponse(bts)
if is {
	response, err := gopherpc.ParseResponse(bts)
	if err != nil {
		log.Fatal(err)
	}
	err = response.ParseResult(new(Data))
    if err != nil {
		log.Fatal(err)
	}
} else {
	response, err := gopherpc.ParseError(bts)
	if err != nil {
		log.Fatal(err)
	}
}
```
