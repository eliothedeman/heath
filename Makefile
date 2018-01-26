
proto:
	protoc -I=. --go_out=$(GOPATH)/src/ block/*.proto

test: proto
	go test -cover ./...
