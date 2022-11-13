# bashpipe

bashpipe is a http server that receive command and execute to local bash. it's used to execute command at remote machine, send command over http request and server will execute it as bash command. 

## Usage
```
$ go run main.go
server started on port 8080

$ curl http://localhost:8080 -d "ls"
README.MD
go.mod
lib
main.go
```