KVSTORE in go
------------------
1. Currently this code supports three basic operations 'set', 'get', and 'delete'.
2. For demo, start listener.go using following command
		go run listener.go
	
	And then run the python file 'demo.py' in another terminal. It will load demo key value pairs from 'input.txt', send the commands alongwith loaded data to listener.

References:
-------------
1. Network Programming with Go - http://jan.newmarch.name/go/
2. In-memory key value store in C, Go and Python - http://www.darkcoding.net/software/in-memory-key-value-store-in-c-go-and-python/
3. Taking Go for a spin - http://pauladamsmith.com/blog/2011/01/go.html
