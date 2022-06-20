# api-go-lang-fiber
Fiber is an Express inspired web framework built on top of Fasthttp


## Install go lang on Ubuntu
check for old version:
```
go version

Command 'go' not found, but can be installed with:

sudo snap install go         # version 1.18.3, or
sudo apt  install golang-go  # version 2:1.13~1ubuntu2
sudo apt  install gccgo-go   # version 2:1.13~1ubuntu2
```
install 
```
sudo apt  install golang-go
```

check version again:
```
go version
// result
go version go1.13.8 linux/amd64

```

in project folder rin command replaced with yours github url (PS: WTF!?):
```sh
go mod init github.com/sivanov/api-go-lang-fiber
// result:
go: creating new go.mod: module github.com/sivanov/api-go-lang-fiber
```

Install Fiber version 2:
```
go get -u github.com/gofiber/fiber/v2
```
If you using VSC editor it is good idea to download this extension for Go lang intelisence:
[https://marketplace.visualstudio.com/items?itemName=golang.Go](https://marketplace.visualstudio.com/items?itemName=golang.Go)


create new file with name server.go (see code from file)

**Notice: need to stop and start server on each file change, Fiber do not have autoreload functionallity**

Adding JSON support and responce

need to instal additional package for JSON format
```
go get github.com/goccy/go-json
```

Status: Error on server start and  bad official documentation:
```
go run server.go 
# command-line-arguments
./server.go:5:2: imported and not used: "github.com/goccy/go-json" as json
```