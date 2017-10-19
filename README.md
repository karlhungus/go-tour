# go-tour


exercises from a tour of go https://tour.golang.org/welcome/1


#### Notes

Static binary, for alpine docker

```shell
go get -u
go build --tags netgo --ldflags '-extldflags "-lm -lstdc++ -static"'
```
