
proto:
	protoc -I=. --go_out=$(GOPATH)/src/ block/*.proto
	protoc -I=. --go_out=$(GOPATH)/src/ wire/*.proto

test: proto
	go test -cover ./...
