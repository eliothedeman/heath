
proto:
	cd block && protoc --go_out=plugins=grpc,import_path=github.com/eliothedeman/heath/block:. *.proto

test: proto
	go test ./...
