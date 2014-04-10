#KVSTORE in go
1. This is a simple key value store implemented in `go`
2. Currently this code supports three basic operations 'set', 'get', and 'delete'.


| Command	| Parameter			| Description  
| -------------	|:-------------:		| -----------
| get		| key `string` 			| returns the value corresponding to passed key
| set		| key `string` value `string`  	| saves the passed key value pair
| delete	| key `string`			| deletes the key value pair identified by the passed key

#Installation and Test

```
go get github.com/sk4x0r/kvstore
go test github.com/sk4x0r/kvstore
```

#Dependencies
Library depends upon ZeroMQ 4, which can be installed from [github.com/pebbe/zmq4](github.com/pebbe/zmq4).

#References:
1. Network Programming with Go - http://jan.newmarch.name/go/
2. In-memory key value store in C, Go and Python - http://www.darkcoding.net/software/in-memory-key-value-store-in-c-go-and-python/
3. Taking Go for a spin - http://pauladamsmith.com/blog/2011/01/go.html
