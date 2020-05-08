.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/license/shared.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/license/space.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/license/key.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/license/certificate.proto
