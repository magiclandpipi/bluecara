```
cd $GOPATH/src/github.com/chipo/proto

protoc -I . \
  -I $GOPATH/src \
  -I $GOPATH/src/github.com/googleapis \
  --go_out=paths=source_relative:. \
  --go-grpc_out=paths=source_relative:. \
  experience.proto
```