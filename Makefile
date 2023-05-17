.PHONY: 
	- gen
	- upgrade

gen: 
	protoc --go_out=. --twirp_out=. rpc/haberdasher/service.proto

upgrade:
	go get -u github.com/twitchtv/twirp/protoc-gen-twirp
	go get -u github.com/golang/protobuf/protoc-gen-go

runserver:
	go run cmd/server/main.go

runclient:
	go run cmd/client/main.go