# Parse Request.

After you have received an array of bytes containing the body, pass them as the first parameter, and as the second parameter, pass a pointer to the object that should be represented by the "parameters" key. If the parsing is successful, then the function will return you only the "request" variable, otherwise it will return the "response" variable already in the format of the error presentation according to the standard JSON-RPC, and will also return an error.

```go
request, response, err := gopherpc.ParseRequest(bts, new(Data))
```

In order to return a response, use the "Response" method of the "Request" object, passing it as an argument a pointer to an object that represents the value by the "result" key. If you want to return an error, use the "Error" method of the "Request" object, passing it an error code as the first argument, and a string describing the error as the second argument. Note that, according to the standard, if the request was a notification, then both methods will return an error.

```go

...

response, err := request.Response(&Some{Hello: "World"})
if err != nil {
	t.Fatal(err)
}

...

response, err := request.Error(gopherpc.ErrMethodNotFound, "Hello world")
if err != nil {
	t.Fatal(err)
}

...

bts, err = response.Marshall()
if err != nil {
	t.Fatal(err)
}
```

# Parse Response.

Before parsing the answer, you can use the "IsResponse()" function to find out if the answer is a normal answer or an error. As in the case of parsing a request, you need to pass an array of bytes containing the body as the first argument, and pass a pointer to an object as the second argument, which can be used to represent the value by the "result" key.

```go

is := gopherpc.IsResponse(bts)
if is {
	response, err := gopherpc.ParseResponse(bts, new(Some))
	if err != nil {
		t.Fatal(err)
	}
} else {
	response, err := gopherpc.ParseError(bts)
	if err != nil {
		t.Fatal(err)
	}
}
```
